package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/peractio/gdk/pkg/env"
	"github.com/peractio/gdk/pkg/errorx/v2"
	"github.com/peractio/gdk/pkg/logx"
	"github.com/peractio/gdk/pkg/resync"
)

var (
	onceNewPGXClient      resync.Once
	onceNewPGXClientRes   *PGXClient
	onceNewPGXClientErr   error
	onceNewPGXClientClose resync.Once
)

// NewPGXClient creates a postgre client using pgx.
func NewPGXClient(ctx context.Context, config *PostgreConfiguration) (*PGXClient, error) {
	onceNewPGXClient.Do(func() {
		const op errorx.Op = "database.NewPGXClient"

		// Parse configuration.
		dbConfig, err := pgxpool.ParseConfig(config.Address)
		if err != nil {
			onceNewPGXClientErr = errorx.E(err, op)
			return
		}
		if !env.IsProduction() {
			dbConfig.ConnConfig.Logger = logx.NewPGX()
		}
		dbConfig.MaxConns = config.MaxConnection
		dbConfig.MinConns = config.MinConnection
		dbConfig.MaxConnLifetime = time.Duration(config.MaxConnectionLifetime) * time.Second
		dbConfig.MaxConnIdleTime = time.Duration(config.MaxConnectionIdleTime) * time.Second

		// Create database connection.
		conn, err := pgxpool.ConnectConfig(ctx, dbConfig)
		if err != nil {
			onceNewPGXClientErr = errorx.E(err, op)
			return
		}

		onceNewPGXClientRes = &PGXClient{
			conn: conn,
		}
	})

	return onceNewPGXClientRes, onceNewPGXClientErr
}

type PGXClient struct {
	conn *pgxpool.Pool
}

// Get returns connection client to postgres.
func (p *PGXClient) Get(_ context.Context) (*pgxpool.Pool, error) {
	const op errorx.Op = "database/PGXClient.Get"

	if p.conn == nil {
		return nil, errorx.E("connection not found", op, errorx.CodeNotFound)
	}

	return p.conn, nil
}

// Close closes all connection to the postgres.
func (p *PGXClient) Close() {
	onceNewPGXClientClose.Do(func() {
		p.conn.Close()
	})
}
