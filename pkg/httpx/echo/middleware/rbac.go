package middleware

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rizalgowandy/gdk/pkg/auth"
	"github.com/rizalgowandy/gdk/pkg/rbac"
)

// RBAC returns a middleware function that performs role-based access control
func RBAC(rbacManager *rbac.Manager, authOperator *auth.Operator) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get JWT claims from context
			claims, ok := c.Get("user").(auth.Claims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"message": "unauthorized",
				})
			}

			// Check permission
			hasPermission, err := rbacManager.Enforce(
				strconv.Itoa(claims.UserID),
				c.Request().URL.Path,
				c.Request().Method,
			)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": "error checking permissions",
				})
			}

			if !hasPermission {
				return c.JSON(http.StatusForbidden, echo.Map{
					"message": "forbidden: insufficient permissions",
				})
			}

			return next(c)
		}
	}
}
