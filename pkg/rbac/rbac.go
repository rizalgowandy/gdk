package rbac

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
)

// Manager represents the role-based access control manager
type Manager struct {
	enable   bool
	enforcer *casbin.Enforcer
	mutex    sync.RWMutex
}

// NewManager creates a new RBAC manager
// It initializes the Casbin enforcer with a model and a policy file.
// The model defines the structure of the policy, and the policy file contains the actual rules.
//
// Example usage:
//
//	policyFile := "pkg/rbac/policy_example.csv"
func NewManager(enable bool, policyFile string) (*Manager, error) {
	// Clean the path to prevent directory traversal attacks
	policyFile = filepath.Clean(policyFile)

	// Load model from string
	m, err := model.NewModelFromString(modelText)
	if err != nil {
		return nil, fmt.Errorf("create model: %w", err)
	}

	// Check if a policy file exists, if not, create it
	if _, err := os.Stat(policyFile); os.IsNotExist(err) {
		file, err := os.Create(policyFile)
		if err != nil {
			return nil, fmt.Errorf("create policy file: %w", err)
		}
		if err := file.Close(); err != nil {
			return nil, fmt.Errorf("close policy file: %w", err)
		}
	}

	// Initialize file adapter
	a := fileadapter.NewAdapter(policyFile)

	// Create enforcer
	enforcer, err := casbin.NewEnforcer(m, a)
	if err != nil {
		return nil, fmt.Errorf("create enforcer: %w", err)
	}

	// Load policies from a file
	if err := enforcer.LoadPolicy(); err != nil {
		return nil, fmt.Errorf("load policy: %w", err)
	}

	return &Manager{
		enable:   enable,
		enforcer: enforcer,
		mutex:    sync.RWMutex{},
	}, nil
}

// AddRoleForUser adds a role for a user
func (r *Manager) AddRoleForUser(user, role string) (bool, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	added, err := r.enforcer.AddRoleForUser(user, role)
	if err != nil {
		return false, fmt.Errorf("add role for user: %w", err)
	}
	return added, nil
}

// DeleteRoleForUser removes a role from a user
func (r *Manager) DeleteRoleForUser(user, role string) (bool, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	deleted, err := r.enforcer.DeleteRoleForUser(user, role)
	if err != nil {
		return false, fmt.Errorf("delete role for user: %w", err)
	}
	return deleted, nil
}

// HasRoleForUser checks if a user has a role
func (r *Manager) HasRoleForUser(user, role string) (bool, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	hasRole, err := r.enforcer.HasRoleForUser(user, role)
	if err != nil {
		return false, fmt.Errorf("check role for user: %w", err)
	}
	return hasRole, nil
}

// GetRolesForUser gets roles for a user
func (r *Manager) GetRolesForUser(user string) ([]string, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	roles, err := r.enforcer.GetRolesForUser(user)
	if err != nil {
		return nil, fmt.Errorf("get roles for user: %w", err)
	}
	return roles, nil
}

// Enforce checks permission for a user
func (r *Manager) Enforce(sub, obj, act string) (bool, error) {
	if !r.enable {
		return true, nil
	}

	r.mutex.RLock()
	defer r.mutex.RUnlock()
	allowed, err := r.enforcer.Enforce(sub, obj, act)
	if err != nil {
		return false, fmt.Errorf("enforce policy: %w", err)
	}
	return allowed, nil
}

// AssignUserRole assigns a role to a user based on the role type
// It should be called after the user login or signup
// Usually on login, so when the database is changes, the next login user can get the new role
func (r *Manager) AssignUserRole(userID int, roleType ...Role) error {
	userIDStr := strconv.Itoa(userID)
	for _, role := range roleType {
		if _, err := r.AddRoleForUser(userIDStr, role.String()); err != nil {
			return err
		}
	}
	return nil
}
