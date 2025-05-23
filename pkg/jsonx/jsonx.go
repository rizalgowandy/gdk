package jsonx

import (
	"io"

	"github.com/rizalgowandy/gdk/pkg/syncx"
)

//go:generate mockgen -destination=jsonx_mock.go -package=jsonx -source=jsonx.go

// OperatorItf interface for json library.
type OperatorItf interface {
	Unmarshal(data []byte, v any) error
	Marshal(v any) ([]byte, error)
	NewEncoder(w io.Writer) EncoderItf
	NewDecoder(r io.Reader) DecoderItf
}

// EncoderItf interface for json library encoder.
type EncoderItf interface {
	Encode(v any) error
}

// DecoderItf interface for json library decoder.
type DecoderItf interface {
	Decode(v any) error
}

var (
	onceNew    syncx.Once
	onceNewRes OperatorItf
)

// New returns a json operator.
func New() OperatorItf {
	onceNew.Do(func() {
		onceNewRes = NewJSONIterator()
	})

	return onceNewRes
}

// Unmarshal copy input data to interface.
func Marshal(v any) ([]byte, error) {
	return New().Marshal(v)
}

// Marshal returns bytes of interface.
func Unmarshal(data []byte, v any) error {
	return New().Unmarshal(data, v)
}

// NewEncoder returns encoder to encode data to buffer.
func NewEncoder(w io.Writer) EncoderItf {
	return New().NewEncoder(w)
}

// NewDecoder returns decoder to decode data to buffer.
func NewDecoder(r io.Reader) DecoderItf {
	return New().NewDecoder(r)
}
