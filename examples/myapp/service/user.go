package service

import (
	"context"

	"github.com/peractio/gdk/examples/myapp"
	"github.com/peractio/gdk/pkg/errorx"
)

type UserService struct {
	repository myapp.UserRepository
}

func (u *UserService) FindUserByID(ctx context.Context, id int) (*myapp.User, error) {
	user, err := u.repository.FindUserByID(ctx, id)
	if errorx.Code(err) == errorx.ENOTFOUND {
		// retry another method of finding our user
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

// CreateUser creates a new user in the system.
// Returns EINVALID if the username is blank or already exists.
// Returns ECONFLICT if the username is already in use.
func (u *UserService) CreateUser(ctx context.Context, user *myapp.User) error {
	// Validate username is non-blank.
	if user.Username == "" {
		return &errorx.Error{Code: errorx.EINVALID, Message: "Username is required."}
	}

	// Verify user does not already exist.
	user, err := u.repository.FindUserByUsername(ctx, user.Username)
	if user != nil {
		return &errorx.Error{
			Code:    errorx.ECONFLICT,
			Message: "Username is already in use. Please choose a different username.",
		}
	}
	if errorx.Code(err) != errorx.ENOTFOUND {
		return &errorx.Error{
			Code:    errorx.EINTERNAL,
			Message: "An internal error has occurred. Please contact technical support.",
		}
	}

	// Create user.
	err = u.repository.CreateUser(ctx, user)
	if err != nil {
		return &errorx.Error{
			Code:    errorx.EINTERNAL,
			Message: "An internal error has occurred. Please contact technical support.",
		}
	}

	return nil
}
