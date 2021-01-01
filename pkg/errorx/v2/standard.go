package errorx

import (
	"fmt"
)

// Recreate the errors.New functionality of the standard Go errors package
// so we can create simple text errors when needed.

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

// Error returns the string representation of the error message.
func (e *errorString) Error() string {
	return e.s
}

// New returns an error that formats as the given text. It is intended to
// be used as the error-typed argument to the E function.
func New(text string) error {
	return &errorString{
		s: text,
	}
}

// Errorf is equivalent to fmt.Errorf, but allows clients to import only this
// package for all error handling.
func Errorf(format string, args ...interface{}) error {
	return &errorString{
		s: fmt.Sprintf(format, args...),
	}
}
