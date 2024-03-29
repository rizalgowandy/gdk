package interceptor

import (
	"context"

	"github.com/rizalgowandy/gdk/pkg/logx"
	"github.com/rizalgowandy/gdk/pkg/queue/nsqx"
)

// RequestID is a middleware that inject request id to the context if it doesn't exist.
func RequestID(
	ctx context.Context,
	consumer *nsqx.Consumer,
	handler nsqx.ConsumerHandler,
) error {
	return handler(logx.ContextWithRequestID(ctx), consumer)
}
