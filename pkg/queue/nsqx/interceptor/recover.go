package interceptor

import (
	"context"
	"runtime/debug"

	"github.com/peractio/gdk/pkg/queue/nsqx"
	"github.com/peractio/gdk/pkg/stack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Recover is a middleware that recovers server from panic.
// Recover also dumps stack trace on panic occurrence.
func Recover() nsqx.ConsumerInterceptor {
	return func(ctx context.Context, consumer *nsqx.Consumer, handler nsqx.ConsumerHandler) (err error) {
		defer func() {
			if err := recover(); err != nil {
				log.WithLevel(zerolog.PanicLevel).
					Interface("err", err).
					Interface("stack", stack.ToArr(stack.Trim(debug.Stack()))).
					Str("topic", consumer.Topic).
					Str("channel", consumer.Channel).
					Msg("recovered")
			}
		}()

		return handler(ctx, consumer)
	}
}
