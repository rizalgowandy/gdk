package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	// DefaultBodyLimitConfig is the default BodyLimit middleware config.
	DefaultBodyLimitConfig = middleware.BodyLimitConfig{
		Skipper: middleware.DefaultSkipper,
		// Limit can be specified as `4x` or `4xB`
		// where x is one of the multiple from K, M, G, T or P.
		Limit: "2M",
	}
)

// BodyLimit returns a BodyLimit middleware.
//
// BodyLimit middleware sets the maximum allowed size for a request body, if the
// size exceeds the configured limit, it sends "413 - Request Entity Too Large"
// response. The BodyLimit is determined based on both `Content-Length` request
// header and actual content read, which makes it super secure.
func BodyLimit() echo.MiddlewareFunc {
	return middleware.BodyLimitWithConfig(DefaultBodyLimitConfig)
}
