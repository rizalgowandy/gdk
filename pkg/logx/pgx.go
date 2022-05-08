package logx

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/rizalgowandy/gdk/pkg/errorx/v2"
	"github.com/rizalgowandy/gdk/pkg/regex"
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

	// Sanitize sql query by replacing \t and \n as space.
	if data != nil {
		val, ok := data["sql"]
		if ok {
			query, ok := val.(string)
			if ok {
				query = strings.ReplaceAll(query, "\t", " ")
				query = strings.ReplaceAll(query, "\n", " ")
				if sanitized, replaceErr := regex.ReplaceAllString(`\s+`, query, " "); replaceErr == nil {
					query = sanitized
				}
				data["sql"] = query
			}
		}
	}

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
