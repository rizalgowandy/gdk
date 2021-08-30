package mux

import (
	"net/http"

	"github.com/rizalgowandy/gdk/pkg/env"
	"github.com/rizalgowandy/gdk/pkg/jsonx"
)

type Writer struct {
	http.ResponseWriter
	StatusCode int
	// Response contains the whole operation response data.
	// Since certain operation has a big response data,
	// response will only be filled on [development, staging] environment.
	Response map[string]interface{}
}

func NewWriter(w http.ResponseWriter) *Writer {
	return &Writer{w, http.StatusOK, map[string]interface{}{}}
}

func (w *Writer) WriteHeader(code int) {
	w.StatusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *Writer) Write(b []byte) (int, error) {
	if !env.IsProduction() {
		_ = jsonx.Unmarshal(b, &w.Response)
	}

	return w.ResponseWriter.Write(b)
}
