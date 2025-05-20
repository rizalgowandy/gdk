package echo

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// By default, Echo does not bind with format id=1,2,3.
// DefaultBinder will bind with format id=1&id=2&id=3.
// CustomBinder is a custom binder for Echo framework that handles query parameters and JSON binding for non-GET requests.
// It supports binding to struct fields, including embedded structs and slices.
// It also provides custom error handling for unsupported types.

type CustomBinder struct {
	echo.DefaultBinder
}

func (cb *CustomBinder) Bind(i any, c echo.Context) error {
	// Handle JSON binding for non-GET requests
	if c.Request().Method != http.MethodGet {
		if strings.HasPrefix(
			c.Request().Header.Get(echo.HeaderContentType),
			echo.MIMEApplicationJSON,
		) {
			if err := cb.DefaultBinder.Bind(i, c); err != nil {
				return fmt.Errorf("binding JSON: %w", err)
			}
		}
	}

	// Validate that we have a pointer to a struct
	val := reflect.ValueOf(i)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return errors.New("binder expects pointer to struct")
	}

	elem := val.Elem()
	if elem.Kind() != reflect.Struct {
		return errors.New("binder expects struct")
	}

	// Parse query parameters for the main struct and any embedded structs
	return cb.bindStruct(elem, c)
}

// bindStruct binds query parameters to struct fields
func (cb *CustomBinder) bindStruct(elem reflect.Value, c echo.Context) error {
	typ := elem.Type()
	for i := range typ.NumField() {
		field := typ.Field(i)
		fieldVal := elem.Field(i)

		// Handle embedded structs
		if field.Anonymous && fieldVal.Kind() == reflect.Struct {
			if err := cb.bindStruct(fieldVal, c); err != nil {
				return err
			}
			continue
		}

		// Get query parameter name from tag
		queryTag := field.Tag.Get("query")
		if queryTag == "" {
			continue
		}

		// Get query parameter value
		raw := c.QueryParam(queryTag)
		if raw == "" {
			continue
		}

		// Check if field can be set
		if !fieldVal.CanSet() {
			continue
		}

		// Handle binding based on field type
		if err := cb.bindValue(fieldVal, field.Type, raw, queryTag); err != nil {
			return err
		}
	}
	return nil
}

// bindValue binds a string value to a reflect.Value based on its type
func (cb *CustomBinder) bindValue(
	fieldVal reflect.Value,
	fieldType reflect.Type,
	raw, queryTag string,
) error {
	// Handle pointer types
	if fieldVal.Kind() == reflect.Ptr {
		elemType := fieldType.Elem()
		newVal := reflect.New(elemType)

		// Bind to the element the pointer points to
		if err := cb.bindValue(newVal.Elem(), elemType, raw, queryTag); err != nil {
			return err
		}

		fieldVal.Set(newVal)
		return nil
	}

	// Handle slice types
	if fieldVal.Kind() == reflect.Slice {
		return cb.bindSlice(fieldVal, fieldType, raw, queryTag)
	}

	// Handle primitive types
	return cb.bindPrimitive(fieldVal, raw, queryTag)
}

// bindSlice binds a comma-separated string to a slice
func (cb *CustomBinder) bindSlice(
	fieldVal reflect.Value,
	fieldType reflect.Type,
	raw, queryTag string,
) error {
	parts := strings.Split(raw, ",")
	elemType := fieldType.Elem()

	// Handle pointer element types in slices
	if elemType.Kind() == reflect.Ptr {
		sliceVal := reflect.MakeSlice(fieldType, len(parts), len(parts))
		for i, part := range parts {
			elemVal := reflect.New(elemType.Elem())
			if err := cb.bindPrimitive(elemVal.Elem(), part, queryTag); err != nil {
				return err
			}
			sliceVal.Index(i).Set(elemVal)
		}
		fieldVal.Set(sliceVal)
		return nil
	}

	// Handle non-pointer element types
	switch elemType.Kind() {
	case reflect.String:
		fieldVal.Set(reflect.ValueOf(parts))
	case reflect.Int, reflect.Int64:
		var vals []int64
		for _, p := range parts {
			v, err := strconv.ParseInt(strings.TrimSpace(p), 10, 64)
			if err != nil {
				return fmt.Errorf("invalid int in '%s': %w", queryTag, err)
			}
			vals = append(vals, v)
		}
		// Convert to the correct slice type (int or int64)
		if elemType.Kind() == reflect.Int {
			ints := make([]int, len(vals))
			for i, v := range vals {
				ints[i] = int(v)
			}
			fieldVal.Set(reflect.ValueOf(ints))
		} else {
			fieldVal.Set(reflect.ValueOf(vals))
		}
	case reflect.Uint:
		var vals []uint
		for _, p := range parts {
			v, err := strconv.ParseUint(strings.TrimSpace(p), 10, 0)
			if err != nil {
				return fmt.Errorf("invalid uint in '%s': %w", queryTag, err)
			}
			vals = append(vals, uint(v))
		}
		fieldVal.Set(reflect.ValueOf(vals))
	case reflect.Float64:
		var vals []float64
		for _, p := range parts {
			v, err := strconv.ParseFloat(strings.TrimSpace(p), 64)
			if err != nil {
				return fmt.Errorf("invalid float64 in '%s': %w", queryTag, err)
			}
			vals = append(vals, v)
		}
		fieldVal.Set(reflect.ValueOf(vals))
	default:
		return fmt.Errorf("unsupported slice element type: %v", elemType.Kind())
	}

	return nil
}

// bindPrimitive binds a string value to a primitive type
func (cb *CustomBinder) bindPrimitive(fieldVal reflect.Value, raw, queryTag string) error {
	trimmed := strings.TrimSpace(raw)

	switch fieldVal.Kind() {
	case reflect.String:
		fieldVal.SetString(raw) // Don't trim spaces for strings by default
	case reflect.Int, reflect.Int64:
		v, err := strconv.ParseInt(trimmed, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid int in '%s': %w", queryTag, err)
		}
		fieldVal.SetInt(v)
	case reflect.Uint:
		v, err := strconv.ParseUint(trimmed, 10, 0)
		if err != nil {
			return fmt.Errorf("invalid uint in '%s': %w", queryTag, err)
		}
		fieldVal.SetUint(v)
	case reflect.Float64:
		v, err := strconv.ParseFloat(trimmed, 64)
		if err != nil {
			return fmt.Errorf("invalid float64 in '%s': %w", queryTag, err)
		}
		fieldVal.SetFloat(v)
	case reflect.Bool:
		v, err := strconv.ParseBool(trimmed)
		if err != nil {
			return fmt.Errorf("invalid bool in '%s': %w", queryTag, err)
		}
		fieldVal.SetBool(v)
	default:
		return fmt.Errorf("unsupported primitive type: %v", fieldVal.Kind())
	}

	return nil
}
