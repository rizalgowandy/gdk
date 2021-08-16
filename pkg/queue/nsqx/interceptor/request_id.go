package interceptor

import (
	"context"

	"github.com/peractio/gdk/pkg/logx"
	"github.com/peractio/gdk/pkg/queue/nsqx"
)

// RequestID is a middleware that inject request id to the context if it doesn't exists.
func RequestID(
	ctx context.Context,
	consumer *nsqx.Consumer,
	handler nsqx.ConsumerHandler,
) error {
	return handler(logx.ContextWithRequestID(ctx), consumer)
}
