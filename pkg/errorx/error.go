package errorx

import (
	"bytes"
	"fmt"
)

// Error defines a standard application error.
type Error struct {
	// Machine-readable error code.
	Code string

	// Human-readable message.
	Message string

	// Logical operation and nested error.
	Op  string
	Err error
}

// Error returns the string representation of the error message.
func (e *Error) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		_, err := fmt.Fprintf(&buf, "%s: ", e.Op)
		if err != nil {
			buf.WriteString("add buffer failure: op exists")
		}
	}

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			_, err := fmt.Fprintf(&buf, "<%s> ", e.Code)
			if err != nil {
				buf.WriteString("add buffer failure: code exists")
			}
		}
		buf.WriteString(e.Message)
	}
	return buf.String()
}
