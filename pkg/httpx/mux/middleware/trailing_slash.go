package middleware

import (
	"net/http"
	"strings"
)

// RemoveTrailingSlash returns a root level (before router) middleware
// which removes a trailing slash from the request URI.
func RemoveTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		url := req.URL
		path := url.Path
		qs := req.URL.RawQuery
		l := len(path) - 1
		if l > 0 && strings.HasSuffix(path, "/") {
			path = path[:l]
			uri := path
			if qs != "" {
				uri += "?" + qs
			}

			// Forward
			req.RequestURI = uri
			url.Path = path
		}
		next.ServeHTTP(res, req)
	})
}
