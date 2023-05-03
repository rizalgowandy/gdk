package interceptor

import (
	"context"
	"fmt"
	"time"

	"github.com/rizalgowandy/gdk/pkg/logx"
	"github.com/rizalgowandy/gdk/pkg/queue/natsx"
	"github.com/rizalgowandy/gdk/pkg/tags"
)

// Logger is a middleware that logs the current result from request.
func Logger(
	ctx context.Context,
	subscriber *natsx.SubscriberImpl,
	handler natsx.SubscriberHandler,
) error {
	start := time.Now()
	err := handler(ctx, subscriber)
	if err != nil {
		logx.ERR(ctx, err, subscriber.Method)
		return err
	}

	logx.DBG(
		ctx,
		logx.KV{tags.Latency: time.Since(start).String()},
		fmt.Sprintf("operation subscriber %s success", subscriber.Method),
	)
	return nil
}
