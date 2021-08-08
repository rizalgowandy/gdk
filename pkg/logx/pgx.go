package logx

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/peractio/gdk/pkg/errorx/v2"
)

func NewPGX() *PGX {
	return &PGX{}
}

type PGX struct{}

func (p *PGX) Log(
	ctx context.Context,
	level pgx.LogLevel,
	msg string,
	data map[string]interface{},
) {
	err := errorx.E("db operation error", errorx.Fields(data))

	switch level {
	case pgx.LogLevelError:
		ERR(ctx, err, msg)
	case pgx.LogLevelWarn:
		WRN(ctx, err, msg)
	case pgx.LogLevelInfo:
		INF(ctx, data, msg)
	case pgx.LogLevelDebug:
		DBG(ctx, data, msg)
	default:
		DBG(ctx, data, msg)
	}
}
