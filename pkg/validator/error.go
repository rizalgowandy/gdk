package validator

import "errors"

var (
	errUnsupportedInputType = errors.New("unsupported input type")
	errInvalidFormatRFC3339 = errors.New("must comply with RFC3339")
)
