package errorx

import (
	"bytes"
)

// Separator is the string used to separate nested errors.
// By default, to make errors easier on the eye, nested errors are
// indented on a new line.
// A server may instead choose to keep each error
// on a single line by modifying the separator string,
// perhaps to ":: " or "-> ".
var Separator = ":\n\t"

// Error defines a standard application error.
type Error struct {
	// Machine-readable error code.
	Code Code `json:"code,omitempty"`

	// Human-readable message.
	Message string `json:"message,omitempty"`

	// Logical operation and nested error.
	Op  Op    `json:"op,omitempty"`
	Err error `json:"-"`
}

// Error returns the string representation of the error message.
//
// Example:
//  userService.FindUserByID: <internal> Internal server error.:
//      accountGateway.FindUserByID: <gateway> Gateway server error.:
//      io.Write: Unknown error.
func (e *Error) Error() string {
	b := new(bytes.Buffer)
	if e.Op != "" {
		pad(b, ": ")
		b.WriteString(string(e.Op))
	}
	if e.Op != "" && (e.Code != Unknown || e.Message != "") {
		pad(b, ": ")
	}
	if e.Code != Unknown {
		b.WriteString("<" + string(e.Code) + ">")
		if e.Message != "" {
			b.WriteString(" ")
		}
	}
	if e.Message != "" {
		b.WriteString(e.Message)
	}
	if e.Err != nil {
		// Indent on new line if we are cascading non-empty errorx.Errors.
		if prevErr, ok := e.Err.(*Error); ok {
			if !prevErr.isZero() {
				pad(b, Separator)
				b.WriteString(e.Err.Error())
			}
		} else {
			pad(b, " => ")
			b.WriteString(e.Err.Error())
		}
	}
	if b.Len() == 0 {
		return "no error"
	}
	return b.String()
}

func (e *Error) isZero() bool {
	return e.Code == Unknown && e.Message == "" && e.Op == "" && e.Err == nil
}

// pad appends str to the buffer if the buffer already has some data.
func pad(b *bytes.Buffer, str string) {
	if b.Len() == 0 {
		return
	}
	b.WriteString(str)
}
