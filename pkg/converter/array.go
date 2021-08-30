package converter

import (
	"github.com/rizalgowandy/gdk/pkg/jsonx"
)

// ArrInt convert any value to []int.
func ArrInt(v interface{}) []int {
	switch v := v.(type) {
	case []int:
		return v

	case []int32:
		result := make([]int, len(v))
		for k, one := range v {
			result[k] = Int(one)
		}
		return result

	case []int64:
		result := make([]int, len(v))
		for k, one := range v {
			result[k] = Int(one)
		}
		return result

	case string:
		var result []int
		err := jsonx.Unmarshal([]byte(v), &result)
		if err != nil {
			return nil
		}
		return result

	case []string:
		result := make([]int, len(v))
		for k, vv := range v {
			result[k] = Int(vv)
		}
		return result

	case [][]byte:
		result := make([]int, len(v))
		for k, vv := range v {
			result[k] = Int(vv)
		}
		return result

	default:
		return nil
	}
}

// ToArrStr convert any value to []string.
func ArrStr(v interface{}) []string {
	switch v := v.(type) {
	case string:
		var result []string
		err := jsonx.Unmarshal([]byte(v), &result)
		if err != nil {
			return nil
		}
		return result

	case [][]byte:
		result := make([]string, len(v))
		for k, vv := range v {
			result[k] = string(vv)
		}
		return result

	default:
		return nil
	}
}

// ToArrInt64 convert any value to []int64.
func ArrInt64(v interface{}) []int64 {
	switch v := v.(type) {
	case []int:
		result := make([]int64, len(v))
		for k, one := range v {
			result[k] = Int64(one)
		}
		return result

	case []int32:
		result := make([]int64, len(v))
		for k, one := range v {
			result[k] = Int64(one)
		}
		return result

	case []int64:
		return v

	case string:
		var result []int64
		err := jsonx.Unmarshal([]byte(v), &result)
		if err != nil {
			return nil
		}
		return result

	case []string:
		result := make([]int64, len(v))
		for k, vv := range v {
			result[k] = Int64(vv)
		}
		return result

	case [][]byte:
		result := make([]int64, len(v))
		for k, vv := range v {
			result[k] = Int64(vv)
		}
		return result

	default:
		return nil
	}
}
