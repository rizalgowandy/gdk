package validator

import (
	"time"
)

// RFC3339 validates input is compliance with RFC3339
func RFC3339(value any) error {
	switch v := value.(type) {
	case string:
		if v == "" {
			return nil
		}

		_, err := time.Parse(time.RFC3339, v)
		if err != nil {
			return errInvalidFormatRFC3339
		}
		return nil

	case time.Time, *time.Time:
		return nil
	}

	return errUnsupportedInputType
}
