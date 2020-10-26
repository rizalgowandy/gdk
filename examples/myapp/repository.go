package myapp

import (
	"context"
)

// UserRepository represents a repository for managing users.
type UserRepository interface {
	// Returns a user by ID.
	FindUserByID(ctx context.Context, id int) (*User, error)

	// Returns a user by username.
	FindUserByUsername(ctx context.Context, username string) (*User, error)

	// Creates a new user.
	CreateUser(ctx context.Context, user *User) error
}
