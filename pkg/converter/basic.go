package converter

import (
	"strconv"
	"strings"

	"github.com/peractio/gdk/pkg/jsonx"
)

// String converts any value to string.
func String(v interface{}) string {
	if v == nil {
		return ""
	}
	switch v := v.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int32:
		return strconv.Itoa(int(v))
	case int64:
		return strconv.FormatInt(v, 10)
	case uint64:
		return strconv.FormatInt(int64(v), 10)
	case bool:
		return strconv.FormatBool(v)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case []uint8:
		return string(v)
	default:
		resultJSON, err := jsonx.Marshal(v)
		if err != nil {
			return ""
		}
		return string(resultJSON)
	}
}

// Bool convert any value to boolean.
func Bool(v interface{}) bool {
	switch v := v.(type) {
	case string:
		str := strings.TrimSpace(v)
		result, err := strconv.ParseBool(str)
		if err != nil {
			return false
		}
		return result
	case int, int32, int64:
		return v != 0
	default:
		return false
	}
}

// Int converts any value to int
func Int(v interface{}) int {
	switch v := v.(type) {
	case string:
		str := strings.TrimSpace(v)
		result, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}
		return result
	case int:
		return v
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case []byte:
		result, err := strconv.Atoi(string(v))
		if err != nil {
			return 0
		}
		return result
	default:
		return 0
	}
}

// Int64 converts any value to int64
func Int64(v interface{}) int64 {
	switch v := v.(type) {
	case string:
		str := strings.TrimSpace(v)
		result, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}
		return int64(result)
	case int:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint64:
		return int64(v)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	case []byte:
		result, err := strconv.Atoi(string(v))
		if err != nil {
			return 0
		}
		return int64(result)
	default:
		return 0
	}
}

// Float64 converts any value to float64
func Float64(v interface{}) float64 {
	switch v := v.(type) {
	case string:
		str := strings.TrimSpace(v)
		result, err := strconv.Atoi(str)
		if err != nil {
			return 0
		}
		return float64(result)
	case int:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint64:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return v
	case []byte:
		result, err := strconv.Atoi(string(v))
		if err != nil {
			return 0
		}
		return float64(result)
	default:
		return 0
	}
}
