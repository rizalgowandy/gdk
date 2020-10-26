package errorx

import (
	"encoding/binary"
)

// MarshalErrorAppend marshals an arbitrary error into a byte slice.
// The result is appended to b, which may be nil.
// It returns the argument slice unchanged if the error is nil.
// If the error is not an *Error, it just records the result of err.Error().
// Otherwise it encodes the full Error struct.
func MarshalErrorAppend(err error, b []byte) []byte {
	if err == nil {
		return b
	}
	if e, ok := err.(*Error); ok {
		// This is an errorx.Error. Mark it as such.
		b = append(b, 'E')
		return e.MarshalAppend(b)
	}
	// Ordinary error.
	b = append(b, 'e')
	b = appendString(b, err.Error())
	return b
}

// MarshalError marshals an arbitrary error and returns the byte slice.
// If the error is nil, it returns nil.
// It returns the argument slice unchanged if the error is nil.
// If the error is not an *Error, it just records the result of err.Error().
// Otherwise it encodes the full Error struct.
func MarshalError(err error) []byte {
	return MarshalErrorAppend(err, nil)
}

// UnmarshalError unmarshal the byte slice into an error value.
// If the slice is nil or empty, it returns nil.
// Otherwise the byte slice must have been created by MarshalError or
// MarshalErrorAppend.
// If the encoded error was of type *Error, the returned error value
// will have that underlying type. Otherwise it will be just a simple
// value that implements the error interface.
func UnmarshalError(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	code := b[0]
	b = b[1:]
	switch code {
	case 'e':
		// Plain error.
		var data []byte
		data, _ = getBytes(b)
		return Str(string(data))
	case 'E':
		// Error value.
		var err Error
		_ = err.UnmarshalBinary(b)
		return &err
	default:
		return Str(string(b))
	}
}

func appendString(b []byte, str string) []byte {
	var tmp [16]byte
	N := binary.PutUvarint(tmp[:], uint64(len(str)))
	b = append(b, tmp[:N]...)
	b = append(b, str...)
	return b
}
