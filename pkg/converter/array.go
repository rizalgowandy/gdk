package converter

import (
	"encoding/json"
)

// ToArrInt convert any value to []int.
func ToArrInt(v interface{}) []int {
	switch v := v.(type) {
	case string:
		var result []int
		err := json.Unmarshal([]byte(v), &result)
		if err != nil {
			return nil
		}
		return result
	case []string:
		var result []int
		for _, vv := range v {
			result = append(result, ToInt(vv))
		}
		return result
	case [][]byte:
		var result []int
		for _, vv := range v {
			result = append(result, ToInt(vv))
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
		var result []string
		for _, vv := range v {
			result = append(result, string(vv))
		}
		return result
	default:
		return nil
	}
}
