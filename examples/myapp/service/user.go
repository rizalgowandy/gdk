package service

import (
	"context"

	"github.com/peractio/gdk/examples/myapp"
	"github.com/peractio/gdk/pkg/errorx"
)

type userService struct {
	repository myapp.UserRepository
}

func (u *userService) FindUserByID(ctx context.Context, id int) (*myapp.User, error) {
	user, err := u.repository.FindUserByID(ctx, id)
	if errorx.Code(err) == errorx.ENotFound {
		// retry another method of finding our user
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

// CreateUser creates a new user in the system.
// Returns EInvalid if the username is blank.
// Returns EConflict if the username is already in use.
func (u *userService) CreateUser(ctx context.Context, user *myapp.User) error {
	const op = "userService.CreateUser"

	// Validate username is non-blank.
	if user.Username == "" {
		return &errorx.Error{
			Code:    errorx.EInvalid,
			Message: "Username is required.",
			Op:      op,
			Err:     nil,
		}
	}

	// Verify user does not already exist.
	inUse, err := u.isUsernameInUse(ctx, user.Username)
	if err != nil {
		return &errorx.Error{
			Code:    errorx.EInternal,
			Message: "An internal error has occurred. Please contact technical support.",
			Op:      op,
			Err:     err,
		}
	}
	if inUse {
		return &errorx.Error{
			Code:    errorx.EConflict,
			Message: "Username is already in use. Please choose a different username.",
			Op:      op,
			Err:     nil,
		}
	}

	// Create user.
	err = u.repository.CreateUser(ctx, user)
	if err != nil {
		return &errorx.Error{
			Code:    errorx.EInternal,
			Message: "An internal error has occurred. Please contact technical support.",
			Op:      op,
			Err:     err,
		}
	}

	return nil
}

func (u *userService) isUsernameInUse(ctx context.Context, username string) (bool, error) {
	const op = "isUsernameInUse"

	user, err := u.repository.FindUserByUsername(ctx, username)
	if errorx.Code(err) == errorx.ENotFound {
		return false, nil
	}

	if err != nil {
		return false, &errorx.Error{
			Code:    errorx.EInternal,
			Message: "An internal error has occurred. Please contact technical support.",
			Op:      op,
			Err:     err,
		}
	}

	return user != nil, nil
}
