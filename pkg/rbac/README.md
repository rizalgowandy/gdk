# RBAC

## Overview
This package provides Role-Based Access Control (RBAC) functionality using Casbin. It allows you to define roles, permissions, and enforce access control in your application.

## Getting Started

### Step 1: Create a Policy File
Create a `policy.csv` file to define your roles and permissions. Below is an example:

```csv
p, superadmin, *, *
p, admin, /api/v1/admin/*, *
p, user, /api/v1/*, *
```

- `p` defines a policy.
- The first column is the role.
- The second column is the resource (e.g., URL path).
- The third column is the action (e.g., HTTP method or `*` for all actions).

### Step 2: Initialize the RBAC Manager
Use the `NewManager` function to initialize the RBAC manager with the policy file.

```go
import (
	"github.com/rizalgowandy/gdk/pkg/rbac"
)

func main() {
	policyFile := "path/to/policy.csv"
	rbacManager, err := rbac.NewManager(policyFile)
	if err != nil {
		panic(err)
	}
	// Use rbacManager in your application
}
```

### Step 3: Use the Middleware

#### Auth Middleware
The `Auth` middleware validates JWT tokens and sets the user claims in the context.

```go
import (
	"github.com/labstack/echo/v4"
	"github.com/rizalgowandy/gdk/pkg/auth"
	"github.com/rizalgowandy/gdk/pkg/httpx/echo/middleware"
)

func main() {
	e := echo.New()
	authOperator := auth.NewOperator("your-secret-key")

	e.Use(middleware.Auth(authOperator))
	// Define your routes
	e.Start(":8080")
}
```

#### RBAC Middleware
The `RBAC` middleware enforces role-based access control based on the user's roles and permissions.

```go
import (
	"github.com/labstack/echo/v4"
	"github.com/rizalgowandy/gdk/pkg/httpx/echo/middleware"
	"github.com/rizalgowandy/gdk/pkg/rbac"
)

func main() {
	e := echo.New()
	rbacManager, _ := rbac.NewManager("path/to/policy.csv")
	authOperator := auth.NewOperator("your-secret-key")

	e.Use(middleware.Auth(authOperator))
	e.Use(middleware.RBAC(rbacManager, authOperator))

	// Define your routes
	e.GET("/api/v1/admin/dashboard", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "Welcome Admin!"})
	})

	e.Start(":8080")
}
```

### Notes
- Ensure the `policy.csv` file is accessible and contains the correct permissions.
- The `Auth` middleware must be used before the `RBAC` middleware to ensure user claims are available for RBAC checks.
