package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rizalgowandy/gdk/pkg/balancer"
	"github.com/rizalgowandy/gdk/pkg/env"
	"github.com/rizalgowandy/gdk/pkg/errorx/v2"
	"github.com/rizalgowandy/gdk/pkg/logx"
	"github.com/rizalgowandy/gdk/pkg/syncx"
	"github.com/rizalgowandy/gdk/pkg/tags"
)

var (
	onceNewPGXClient      syncx.Once
	onceNewPGXClientRes   *PGXClient
	onceNewPGXClientErr   error
	onceNewPGXClientClose syncx.Once
)

// NewPGXClient creates a postgre client using pgx.
// First config parameter will always be used for writer.
// Second config and next will always be used for reader.
// If only one configuration is being passed,
// that configuration will be used for both writer and reader.
func NewPGXClient(ctx context.Context, config ...*PostgreConfiguration) (*PGXClient, error) {
	onceNewPGXClient.Do(func() {
		if len(config) == 0 {
			onceNewPGXClientErr = errorx.E("no configuration passed")
			return
		}

		// Create writer.
		writer, err := connect(ctx, &configuration{
			Address:               config[0].Address,
			MinConnection:         config[0].MinConnection,
			MaxConnection:         config[0].MaxConnection,
			MaxConnectionLifetime: config[0].MaxConnectionLifetime,
			MaxConnectionIdleTime: config[0].MaxConnectionIdleTime,
		})
		if err != nil {
			onceNewPGXClientRes = nil
			onceNewPGXClientErr = errorx.E(err, errorx.Fields{tags.Type: "master"})
			return
		}

		// Create reader.
		var readers []*pgxpool.Pool
		for i := 1; i < len(config); i++ {
			reader, errReader := connect(ctx, &configuration{
				Address:               config[i].Address,
				MinConnection:         config[i].MinConnection,
				MaxConnection:         config[i].MaxConnection,
				MaxConnectionLifetime: config[i].MaxConnectionLifetime,
				MaxConnectionIdleTime: config[i].MaxConnectionIdleTime,
			})
			if errReader != nil {
				onceNewPGXClientRes = nil
				onceNewPGXClientErr = errorx.E(errReader, errorx.Fields{tags.Type: "reader", tags.Index: i})
				return
			}
			readers = append(readers, reader)
		}
		if len(readers) == 0 {
			readers = append(readers, writer)
		}
		readerBalancer, err := balancer.NewRoundRobin(toArrItf(readers))
		if err != nil {
			onceNewPGXClientRes = nil
			onceNewPGXClientErr = errorx.E(err, errorx.Fields{tags.Type: "reader balancer"})
			return
		}

		onceNewPGXClientRes = &PGXClient{
			writer:         writer,
			readers:        readers,
			readerBalancer: readerBalancer,
		}
	})

	return onceNewPGXClientRes, onceNewPGXClientErr
}

type PGXClient struct {
	writer         *pgxpool.Pool
	readers        []*pgxpool.Pool
	readerBalancer *balancer.RoundRobin
}

// Get returns connection client to postgres.
func (p *PGXClient) Get(ctx context.Context) (*pgxpool.Pool, error) {
	return p.GetWriter(ctx)
}

func (p *PGXClient) GetWriter(_ context.Context) (*pgxpool.Pool, error) {
	if p.writer == nil {
		return nil, errorx.E("connection not found", errorx.CodeNotFound)
	}

	return p.writer, nil
}

func (p *PGXClient) GetReader(_ context.Context) (*pgxpool.Pool, error) {
	if p.readerBalancer == nil {
		return nil, errorx.E("connection not found", errorx.CodeNotFound)
	}

	reader, ok := p.readerBalancer.Next().(*pgxpool.Pool)
	if !ok {
		return nil, errorx.E("wrong item type in balancer", errorx.CodeInvalid)
	}

	return reader, nil
}

// Close closes all connection to the postgres.
func (p *PGXClient) Close() {
	onceNewPGXClientClose.Do(func() {
		p.writer.Close()
		for k := range p.readers {
			p.readers[k].Close()
		}
	})
}

type configuration struct {
	Address               string
	MinConnection         int32
	MaxConnection         int32
	MaxConnectionLifetime int64
	MaxConnectionIdleTime int64
}

func connect(ctx context.Context, config *configuration) (*pgxpool.Pool, error) {
	// Parse configuration.
	dbConfig, err := pgxpool.ParseConfig(config.Address)
	if err != nil {
		return nil, errorx.E(err)
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
		return nil, errorx.E(err)
	}

	return conn, nil
}

func toArrItf(v []*pgxpool.Pool) []interface{} {
	result := make([]interface{}, len(v))
	for k := range v {
		result[k] = v[k]
	}
	return result
}
