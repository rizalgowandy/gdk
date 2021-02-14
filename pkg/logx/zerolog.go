package logx

import (
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/peractio/gdk/pkg/filex"
	"github.com/peractio/gdk/pkg/tags"
	"github.com/rs/zerolog"
)

var zerologCallerSkip = 2 + zerolog.CallerSkipFrameCount

type ZeroLogger struct {
	client zerolog.Logger
}

func NewZerolog(config *Config) (*ZeroLogger, error) {
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

	// Override caller.
	zerolog.CallerMarshalFunc = func(file string, line int) string {
		file = file[strings.Index(file, config.AppName)+len(config.AppName):]
		return file + ":" + strconv.Itoa(line)
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
	fields map[string]interface{},
	message string,
) {
	l.client.Trace().
		Str(tags.RequestID, requestID).
		Fields(fields).
		Msg(message)
}

func (z *ZeroLogger) Debug(
	requestID string,
	fields map[string]interface{},
	message string,
) {
	l.client.Debug().
		Str(tags.RequestID, requestID).
		Fields(fields).
		Msg(message)
}

func (z *ZeroLogger) Info(
	requestID string,
	fields map[string]interface{},
	message string,
) {
	l.client.Info().
		Str(tags.RequestID, requestID).
		Fields(fields).
		Msg(message)
}

func (z *ZeroLogger) Warn(
	requestID string,
	err error,
	fields map[string]interface{},
	message string,
) {
	l.client.Warn().
		Str(tags.RequestID, requestID).
		Fields(fields).
		Err(err).
		Msg(message)
}

func (z *ZeroLogger) Error(
	requestID string,
	err error,
	fields map[string]interface{},
	message string,
) {
	l.client.Error().
		Str(tags.RequestID, requestID).
		Fields(fields).
		Err(err).
		Msg(message)
}

func (z *ZeroLogger) Fatal(
	requestID string,
	err error,
	fields map[string]interface{},
	message string,
) {
	l.client.Fatal().
		Str(tags.RequestID, requestID).
		Fields(fields).
		Err(err).
		Msg(message)
}
