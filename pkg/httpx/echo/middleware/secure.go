package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Secure returns a Secure middleware.
// Secure middleware provides protection against cross-site scripting (XSS) attack,
// content type sniffing, click-jacking, insecure connection and other code injection
// attacks.
func Secure() echo.MiddlewareFunc {
	return middleware.Secure()
}
