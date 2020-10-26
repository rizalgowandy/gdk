package postgres

import (
	"context"
	"database/sql"

	"github.com/peractio/gdk/examples/myapp"
	"github.com/peractio/gdk/pkg/errorx"
)

type userPostgres struct {
	db sql.DB
}

// FindUserByID returns a user by ID. Returns ENOTFOUND if user does not exist.
func (u *userPostgres) FindUserByID(ctx context.Context, id int) (*myapp.User, error) {
	const op = "userPostgres.FindUserByID"

	query := `
		SELECT id, username
		FROM users
		WHERE id = $1
	`

	var user myapp.User
	err := u.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
	)

	if err == sql.ErrNoRows {
		return nil, &errorx.Error{
			Code: errorx.ENOTFOUND,
			Op:   op,
			Err:  err,
		}
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// CreateUser creates a new user in the system with a default role.
func (u *userPostgres) CreateUser(ctx context.Context, user *myapp.User) error {
	const op = "userPostgres.CreateUser"

	// Perform validation...

	// Insert user record.
	if err := u.insertUser(ctx, user); err != nil {
		return &errorx.Error{Op: op, Err: err}
	}

	// Insert default role.
	if err := u.attachRole(ctx, user.ID, "default"); err != nil {
		return &errorx.Error{Op: op, Err: err}
	}
	return nil
}

// insertUser inserts the user into the database.
func (u *userPostgres) insertUser(ctx context.Context, user *myapp.User) error {
	const op = "insertUser"

	query := `
		INSERT INTO users
	`

	if _, err := u.db.ExecContext(ctx, query, user.ID, user.Username); err != nil {
		return &errorx.Error{Op: op, Err: err}
	}
	return nil
}

// attachRole inserts a role record for a user in the database
func (u *userPostgres) attachRole(ctx context.Context, id int, role string) error {
	const op = "attachRole"

	query := `
		INSERT roles
	`

	if _, err := u.db.ExecContext(ctx, query, id, role); err != nil {
		return &errorx.Error{Op: op, Err: err}
	}
	return nil
}
