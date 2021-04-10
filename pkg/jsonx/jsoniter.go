package jsonx

import (
	"bytes"

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

// Encode copy data to buffer.
func (j *JSONIterator) Encode(buffer *bytes.Buffer, data interface{}) error {
	return j.json.NewEncoder(buffer).Encode(data)
}
