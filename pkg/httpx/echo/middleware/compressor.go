package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	// DefaultGzipConfig is the default Gzip middleware config.
	DefaultGzipConfig = middleware.GzipConfig{
		// skip middleware for:
		// * path for swagger
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "swagger") ||
				strings.Contains(c.Path(), "metrics")
		},
		Level: 5,
	}
)

// Gzip returns a middleware which compresses HTTP response
// using gzip compression scheme.
func Gzip() echo.MiddlewareFunc {
	return middleware.GzipWithConfig(DefaultGzipConfig)
}
