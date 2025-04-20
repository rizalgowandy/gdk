package mux

import (
	"bytes"
	"io"
	"net/http"

	"github.com/rizalgowandy/gdk/pkg/jsonx"
)

// CopyBody returns query param for get method, and body for others.
func CopyBody(r *http.Request) (interface{}, error) {
	if r.Method == http.MethodGet {
		return r.URL.Query(), nil
	}

	// Copy request body to restore it for next request.
	buf, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	// Re-append request body for next request.
	r.Body = io.NopCloser(bytes.NewBuffer(buf))

	// Certain method doesn't have any body, e.g. DELETE.
	if len(buf) == 0 {
		return "", nil
	}

	// Unmarshal binary to struct.
	var res interface{}
	err = jsonx.Unmarshal(buf, &res)
	if err != nil {
		return nil, NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	return res, nil
}
