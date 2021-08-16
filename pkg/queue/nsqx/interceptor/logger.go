package interceptor

import (
	"context"
	"time"

	"github.com/peractio/gdk/pkg/queue/nsqx"
	"github.com/rs/zerolog/log"
)

// Logger is a middleware that logs the current result from request.
func Logger(
	ctx context.Context,
	consumer *nsqx.Consumer,
	handler nsqx.ConsumerHandler,
) error {
	now := time.Now()
	log.Info().
		Str("topic", consumer.Topic).
		Str("channel", consumer.Channel).
		Str("start_time", now.String()).
		Msg("starting a consumer")

	err := handler(ctx, consumer)
	if err != nil {
		log.Error().
			Str("topic", consumer.Topic).
			Str("channel", consumer.Channel).
			Str("latency", time.Since(now).String()).
			Str("finish_time", time.Now().String()).
			Err(err).
			Msg("consumer has finished with error")
		return err
	}

	log.Info().
		Str("topic", consumer.Topic).
		Str("channel", consumer.Channel).
		Str("latency", time.Since(now).String()).
		Str("finish_time", time.Now().String()).
		Msg("consumer has finished")
	return nil
}
