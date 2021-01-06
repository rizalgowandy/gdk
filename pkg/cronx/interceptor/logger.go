package interceptor

import (
	"context"
	"time"

	"github.com/peractio/gdk/pkg/cronx"
	"github.com/rs/zerolog/log"
)

// Logger is a middleware that logs the current job start and finish.
func Logger() cronx.Interceptor {
	return func(ctx context.Context, job *cronx.Job, handler cronx.Handler) error {
		now := time.Now()
		log.Info().
			Str("job", job.Name).
			Str("start_time", now.String()).
			Msg("starting a job")

		err := handler(ctx, job)
		if err != nil {
			log.Error().
				Str("job", job.Name).
				Str("latency", time.Since(now).String()).
				Str("finish_time", time.Now().String()).
				Err(err).
				Msg("job has finished with error")
			return err
		}

		log.Info().
			Str("job", job.Name).
			Str("latency", time.Since(now).String()).
			Str("finish_time", time.Now().String()).
			Msg("job has finished")
		return nil
	}
}
