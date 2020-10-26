package postgres

import (
	"context"
	"database/sql"

	"github.com/peractio/gdk/examples/myapp"
	"github.com/peractio/gdk/pkg/errorx"
)

type UserPostgres struct {
	sql.DB
}

// FindUserByID returns a user by ID. Returns ENOTFOUND if user does not exist.
func (s *UserPostgres) FindUserByID(ctx context.Context, id int) (*myapp.User, error) {
	var user myapp.User

	query := `
		SELECT id, username
		FROM users
		WHERE id = $1
	`

	err := s.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
	)

	if err == sql.ErrNoRows {
		return nil, &errorx.Error{
			Code: errorx.ENOTFOUND,
		}
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
