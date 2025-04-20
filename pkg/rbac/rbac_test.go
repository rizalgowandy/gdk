package rbac

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRBACManager(t *testing.T) {
	// Initialize RBAC Manager
	manager, err := NewManager("policy_example.csv")
	require.NoError(t, err)

	// Test AddRoleForUser
	added, err := manager.AddRoleForUser("1", "admin")
	require.NoError(t, err)
	assert.True(t, added)

	// Test GetRolesForUser
	roles, err := manager.GetRolesForUser("1")
	require.NoError(t, err)
	assert.Contains(t, roles, "admin")

	// Test HasRoleForUser
	hasRole, err := manager.HasRoleForUser("1", "admin")
	require.NoError(t, err)
	assert.True(t, hasRole)

	// Test DeleteRoleForUser
	deleted, err := manager.DeleteRoleForUser("1", "admin")
	require.NoError(t, err)
	assert.True(t, deleted)

	// Test Enforce
	allowed, err := manager.Enforce("1", "/api/v1/resource", "GET")
	require.NoError(t, err)
	assert.False(t, allowed)
}
