package jsonx

import (
	"bytes"

	"github.com/peractio/gdk/pkg/resync"
)

//go:generate mockgen -destination=jsonx_mock.go -package=jsonx -source=jsonx.go

// OperatorItf interface for json library.
type OperatorItf interface {
	Unmarshal(data []byte, v interface{}) error
	Marshal(v interface{}) ([]byte, error)
	Encode(buffer *bytes.Buffer, data interface{}) error
}

var (
	onceNew    resync.Once
	onceNewRes OperatorItf
)

// New returns a json operator.
func New() OperatorItf {
	onceNew.Do(func() {
		onceNewRes = NewJSONIterator()
	})

	return onceNewRes
}
