package jsonx

import (
	"bytes"
	"io"

	jsoniter "github.com/json-iterator/go"
)

// JSONIterator is a json operator using json-iterator library.
type JSONIterator struct {
	json jsoniter.API
}

// NewJSONIterator constructs new json operator using json-iterator library.
func NewJSONIterator() *JSONIterator {
	return &JSONIterator{
		json: jsoniter.ConfigCompatibleWithStandardLibrary,
	}
}

// Unmarshal copy input data to interface.
func (j *JSONIterator) Unmarshal(data []byte, v interface{}) error {
	return j.json.Unmarshal(data, v)
}

// Marshal returns bytes of interface.
func (j *JSONIterator) Marshal(v interface{}) ([]byte, error) {
	return j.json.Marshal(v)
}

// NewEncoder returns encoder to encode data to buffer.
func (j *JSONIterator) NewEncoder(buffer *bytes.Buffer) EncoderItf {
	return j.json.NewEncoder(buffer)
}

// NewDecoder returns decoder to decode data to buffer.
func (j *JSONIterator) NewDecoder(r io.Reader) DecoderItf {
	return j.json.NewDecoder(r)
}
