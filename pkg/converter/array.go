package converter

import (
	"encoding/json"
)

// ToArrInt convert any value to []int.
func ToArrInt(v interface{}) []int {
	switch v := v.(type) {
	case []int:
		return v

	case []int32:
		result := make([]int, len(v))
		for k, one := range v {
			result[k] = ToInt(one)
		}
		return result

	case []int64:
		result := make([]int, len(v))
		for k, one := range v {
			result[k] = ToInt(one)
		}
		return result

	case string:
		var result []int
		err := json.Unmarshal([]byte(v), &result)
		if err != nil {
			return nil
		}
		return result

	case []string:
		result := make([]int, len(v))
		for k, vv := range v {
			result[k] = ToInt(vv)
		}
		return result

	case [][]byte:
		result := make([]int, len(v))
		for k, vv := range v {
			result[k] = ToInt(vv)
		}
		return result

	default:
		return nil
	}
}

// ToArrStr convert any value to []string.
func ToArrStr(v interface{}) []string {
	switch v := v.(type) {
	case string:
		var result []string
		err := json.Unmarshal([]byte(v), &result)
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
