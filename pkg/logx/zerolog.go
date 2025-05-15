package logx

import (
	"io"
	"os"

	"github.com/rizalgowandy/gdk/pkg/filex"
	"github.com/rizalgowandy/gdk/pkg/tags"
	"github.com/rs/zerolog"
)

var zerologCallerSkip = 2 + zerolog.CallerSkipFrameCount

type ZeroLogger struct {
	client zerolog.Logger
}

func NewZerolog(config Config) (*ZeroLogger, error) {
	// Default writer to stderr.
	writer := io.Writer(os.Stderr)

	// Create log file.
	if config.Filename != "" {
		file, err := filex.OpenFile(config.Filename)
		if err != nil {
			return nil, err
		}

		writer = file
	}

	// On debug use pretty print.
	if config.Debug {
		writer = zerolog.ConsoleWriter{Out: writer}
	}

	// Create client.
	client := zerolog.New(writer).
		With().
		Timestamp().
		CallerWithSkipFrameCount(zerologCallerSkip).
		Logger().
		Level(zerolog.TraceLevel)

	// Except development add app name.
	if !config.Debug {
		client = client.With().
			Str(tags.App, config.AppName).
			Logger()
	}

	return &ZeroLogger{client: client}, nil
}

func (z *ZeroLogger) Trace(
	requestID string,
	actorID int,
	fields map[string]any,
	message string,
) {
	event := z.client.Trace().
		Str(tags.RequestID, requestID).
		Fields(fields)

	if actorID != 0 {
		event = event.Int(tags.ActorID, actorID)
	}

	event.Msg(message)
}

func (z *ZeroLogger) Debug(
	requestID string,
	actorID int,
	fields map[string]any,
	message string,
) {
	event := z.client.Debug().
		Str(tags.RequestID, requestID).
		Fields(fields)

	if actorID != 0 {
		event = event.Int(tags.ActorID, actorID)
	}

	event.Msg(message)
}

func (z *ZeroLogger) Info(
	requestID string,
	actorID int,
	fields map[string]any,
	message string,
) {
	event := z.client.Info().
		Str(tags.RequestID, requestID).
		Fields(fields)

	if actorID != 0 {
		event = event.Int(tags.ActorID, actorID)
	}

	event.Msg(message)
}

func (z *ZeroLogger) Warn(
	requestID string,
	actorID int,
	err error,
	fields map[string]any,
	message string,
) {
	event := z.client.Warn().
		Str(tags.RequestID, requestID).
		Fields(fields).
		Err(err)

	if actorID != 0 {
		event = event.Int(tags.ActorID, actorID)
	}

	event.Msg(message)
}

func (z *ZeroLogger) Error(
	requestID string,
	actorID int,
	err error,
	fields map[string]any,
	message string,
) {
	event := z.client.Error().
		Str(tags.RequestID, requestID).
		Fields(fields).
		Err(err)

	if actorID != 0 {
		event = event.Int(tags.ActorID, actorID)
	}

	event.Msg(message)
}

func (z *ZeroLogger) Fatal(
	requestID string,
	actorID int,
	err error,
	fields map[string]any,
	message string,
) {
	event := z.client.Fatal().
		Str(tags.RequestID, requestID).
		Fields(fields).
		Err(err)

	if actorID != 0 {
		event = event.Int(tags.ActorID, actorID)
	}

	event.Msg(message)
}
