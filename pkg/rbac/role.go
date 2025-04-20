package rbac

type Role string

func (role Role) String() string {
	return string(role)
}

// Role types
const (
	RoleSuperAdmin Role = "super_admin"
	RoleAdmin      Role = "admin"
	RoleUser       Role = "user"
)
