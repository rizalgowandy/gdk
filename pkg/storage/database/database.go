package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

//go:generate mockgen -destination=database_mock.go -package=database -source=database.go

// PostgreClientItf is a client to interact with PostgreSQL database using pgx.
type PostgreClientItf interface {
	// Get returns connection client to postgres.
	Get(ctx context.Context) (*pgxpool.Pool, error)

	// GetWriter returns connection client to postgres only for writing.
	GetWriter(ctx context.Context) (*pgxpool.Pool, error)

	// GetReader returns connection client to postgres only for reading.
	GetReader(ctx context.Context) (*pgxpool.Pool, error)

	// Close closes all connection to the postgres.
	Close()
}
