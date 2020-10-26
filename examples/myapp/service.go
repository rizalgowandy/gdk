package myapp

import (
	"context"
)

// UserService represents a service for managing users.
type UserService interface {
	// Returns a user by ID.
	FindUserByID(ctx context.Context, id int) (*User, error)

	// Creates a new user.
	CreateUser(ctx context.Context, user *User) error
}
