package errorx

import (
	"bytes"
	"encoding/binary"
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
	Code Code

	// Human-readable message.
	Message string

	// Logical operation and nested error.
	Op  Op
	Err error
}

// Error returns the string representation of the error message.
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

// MarshalAppend marshals err into a byte slice. The result is appended to b,
// which may be nil.
// It returns the argument slice unchanged if the error is nil.
func (e *Error) MarshalAppend(b []byte) []byte {
	if e == nil {
		return b
	}
	b = appendString(b, e.Message)
	b = appendString(b, string(e.Op))
	b = appendString(b, string(e.Code))
	b = MarshalErrorAppend(e.Err, b)
	return b
}

// MarshalBinary marshals its receiver into a byte slice, which it returns.
// It returns nil if the error is nil. The returned error is always nil.
func (e *Error) MarshalBinary() ([]byte, error) {
	return e.MarshalAppend(nil), nil
}

// UnmarshalBinary unmarshal the byte slice into the receiver, which must be non-nil.
// The returned error is always nil.
func (e *Error) UnmarshalBinary(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	data, b := getBytes(b)
	if data != nil {
		e.Message = string(data)
	}
	data, b = getBytes(b)
	if data != nil {
		e.Op = Op(data)
	}
	data, b = getBytes(b)
	if data != nil {
		e.Code = Code(data)
	}
	e.Err = UnmarshalError(b)
	return nil
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

// getBytes unmarshal the byte slice at b (uvarint count followed by bytes)
// and returns the slice followed by the remaining bytes.
// If there is insufficient data, both return values will be nil.
func getBytes(b []byte) (data, remaining []byte) {
	u, N := binary.Uvarint(b)
	if len(b) < N+int(u) {
		return nil, nil
	}
	if N == 0 {
		return nil, b
	}
	return b[N : N+int(u)], b[N+int(u):]
}
