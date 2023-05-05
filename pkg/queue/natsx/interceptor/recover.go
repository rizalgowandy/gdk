package interceptor

import (
	"context"
	"runtime/debug"

	"github.com/rizalgowandy/gdk/pkg/queue/natsx"
	"github.com/rizalgowandy/gdk/pkg/stack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Recover is a middleware that recovers server from panic.
// Recover also dumps stack trace on panic occurrence.
func Recover() natsx.SubscriberInterceptor {
	return func(ctx context.Context, subscriber *natsx.Subscriber, handler natsx.SubscriberHandler) (err error) {
		defer func() {
			if err := recover(); err != nil {
				log.WithLevel(zerolog.PanicLevel).
					Interface("err", err).
					Interface("stack", stack.ToArr(stack.Trim(debug.Stack()))).
					Str("subject", subscriber.Subject).
					Str("queue", subscriber.Queue).
					Msg("recovered")
			}
		}()

		return handler(ctx, subscriber)
	}
}
