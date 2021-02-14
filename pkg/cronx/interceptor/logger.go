package interceptor

import (
	"context"
	"fmt"
	"time"

	"github.com/peractio/gdk/pkg/cronx"
	"github.com/peractio/gdk/pkg/logx"
	"github.com/peractio/gdk/pkg/tags"
)

// Logger is a middleware that logs the current job start and finish.
func Logger() cronx.Interceptor {
	return func(ctx context.Context, job *cronx.Job, handler cronx.Handler) error {
		start := time.Now()
		err := handler(ctx, job)
		if err != nil {
			logx.ERR(ctx, err, job.Name)
		} else {
			logx.DBG(
				ctx,
				map[string]string{tags.Latency: time.Since(start).String()},
				fmt.Sprintf("Operation cron %s success", job.Name),
			)
		}
		return nil
	}
}
