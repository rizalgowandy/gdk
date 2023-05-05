package interceptor

import (
	"context"

	"github.com/rizalgowandy/gdk/pkg/logx"
	"github.com/rizalgowandy/gdk/pkg/queue/natsx"
)

// RequestID is a middleware that inject request id to the context if it doesn't exist.
func RequestID(
	ctx context.Context,
	subscriber *natsx.Subscriber,
	handler natsx.SubscriberHandler,
) error {
	return handler(logx.ContextWithRequestID(ctx), subscriber)
}
