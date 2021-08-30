package interceptor

import (
	"context"
	"fmt"
	"time"

	"github.com/rizalgowandy/gdk/pkg/logx"
	"github.com/rizalgowandy/gdk/pkg/queue/nsqx"
	"github.com/rizalgowandy/gdk/pkg/tags"
)

// Logger is a middleware that logs the current result from request.
func Logger(
	ctx context.Context,
	consumer *nsqx.Consumer,
	handler nsqx.ConsumerHandler,
) error {
	start := time.Now()
	err := handler(ctx, consumer)
	if err != nil {
		logx.ERR(ctx, err, consumer.Method)
		return err
	}

	logx.DBG(
		ctx,
		logx.KV{tags.Latency: time.Since(start).String()},
		fmt.Sprintf("operation consumer %s success", consumer.Method),
	)
	return nil
}
