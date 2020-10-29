package internal

import (
	"context"

	"github.com/peractio/gdk/examples/myapp/internal/entity"
)

// UserService represents a service for managing users.
type UserService interface {
	// Returns a user by ID.
	FindUserByID(ctx context.Context, id int) (*entity.User, error)

	// Creates a new user.
	CreateUser(ctx context.Context, user *entity.User) error
}
