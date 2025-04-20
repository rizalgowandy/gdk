package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rizalgowandy/gdk/pkg/auth"
	"github.com/rizalgowandy/gdk/pkg/logx"
)

// Auth creates a middleware that validates JWT tokens
func Auth(h *auth.Operator) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "Authorization header missing",
				})
			}

			// Check if it's a Bearer token
			if !strings.HasPrefix(authHeader, "Bearer ") {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "Invalid token format",
				})
			}

			// Extract token
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			// Validate token
			claims, err := h.ValidateToken(tokenString)
			if err != nil {
				logx.ERR(c.Request().Context(), err, "failed to validate token")
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "Invalid or expired token",
				})
			}

			// Set user in context for handlers to use
			c.Set("user", claims)

			// Continue to the next handler
			return next(c)
		}
	}
}
