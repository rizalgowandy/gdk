package converter

import (
	"encoding/json"
	"strconv"
	"strings"
)

// ToString converts any value to string.
func ToString(v interface{}) string {
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
	case bool:
		return strconv.FormatBool(v)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case []uint8:
		return string(v)
	default:
		resultJSON, err := json.Marshal(v)
		if err != nil {
			return ""
		}
		return string(resultJSON)
	}
}

// ToBool convert any value to boolean.
func ToBool(v interface{}) bool {
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

// ToInt converts any value to int
func ToInt(v interface{}) int {
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

// ToInt64 converts any value to int64
func ToInt64(v interface{}) int64 {
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
