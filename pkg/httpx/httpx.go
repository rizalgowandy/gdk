package httpx

import "net/http"

//go:generate mockgen -destination=httpx_mock.go -package=httpx -source=httpx.go

// ClientItf is a http client.
type ClientItf interface {
	// Do sends http request and returns the response.
	Do(req *http.Request) (*http.Response, error)
}

// ReadCloserItf is interface for io.ReadCloser.
type ReadCloserItf interface {
	Read([]byte) (n int, err error)
	Close() error
}
