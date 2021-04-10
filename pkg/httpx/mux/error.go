package mux

import (
	"fmt"
	"net/http"
)

var ErrUnsupportedMediaType = NewHTTPError(http.StatusUnsupportedMediaType)

// NewHTTPError creates a new HTTPError instance.
func NewHTTPError(code int, message ...interface{}) *HTTPError {
	he := &HTTPError{Code: code, Message: http.StatusText(code)}
	if len(message) > 0 {
		he.Message = message[0]
	}
	return he
}

// HTTPError represents an error that occurred while handling a request.
type HTTPError struct {
	Code     int         `json:"-"`
	Message  interface{} `json:"message"`
	Internal error       `json:"-"` // Stores the error returned by an external dependency
}

// Error makes it compatible with `error` interface.
func (he *HTTPError) Error() string {
	if he.Internal == nil {
		return fmt.Sprintf("code=%d, message=%v", he.Code, he.Message)
	}
	return fmt.Sprintf("code=%d, message=%v, internal=%v", he.Code, he.Message, he.Internal)
}

// SetInternal sets error to HTTPError.Internal
func (he *HTTPError) SetInternal(err error) *HTTPError {
	he.Internal = err
	return he
}

// Unwrap satisfies the Go 1.13 error wrapper interface.
func (he *HTTPError) Unwrap() error {
	return he.Internal
}
