package middleware

import (
	"net/http"
)

type (
	// Skipper defines a function to skip middleware.
	// Returning true skips processing the middleware.
	Skipper func(res http.ResponseWriter, req *http.Request) bool
)

// DefaultSkipper returns false which processes the middleware.
func DefaultSkipper(res http.ResponseWriter, req *http.Request) bool {
	return false
}
