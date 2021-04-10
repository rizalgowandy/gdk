package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/peractio/gdk/pkg/env"
)

// Logger returns a middleware that logs HTTP requests
func Logger() echo.MiddlewareFunc {
	// for development, use non json format.
	if env.IsDevelopment() {
		return middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "\n" +
				"time:${time_rfc3339} " +
				"id:${id} " +
				"method:${method} " +
				"uri:${uri} " +
				"status:${status} " +
				"latency_human:${latency_human}",
		})
	}

	// for [staging,beta,uat,production] use json format.
	return middleware.Logger()
}
