package logx

import (
	"context"
	"errors"
)

// KV is wrapper for map[string]interface{}.
type KV map[string]interface{}

var (
	l *ZeroLogger
)

type Config struct {
	Debug    bool
	AppName  string
	Filename string
}

func New(config *Config) (*ZeroLogger, error) {
	if config == nil {
		return nil, errors.New("missing configuration")
	}

	logger, err := NewZerolog(config)
	if err != nil {
		return nil, err
	}

	l = logger
	return l, nil
}

func TRC(ctx context.Context, metadata interface{}, message string) {
	l.Trace(GetRequestID(ctx), Metadata(metadata), message)
}

func DBG(ctx context.Context, metadata interface{}, message string) {
	l.Debug(GetRequestID(ctx), Metadata(metadata), message)
}

func INF(ctx context.Context, metadata interface{}, message string) {
	l.Info(GetRequestID(ctx), Metadata(metadata), message)
}

func WRN(ctx context.Context, err error, message string) {
	l.Warn(GetRequestID(ctx), err, ErrMetadata(err), message)
}

func ERR(ctx context.Context, err error, message string) {
	l.Error(GetRequestID(ctx), err, ErrMetadata(err), message)
}

func FTL(ctx context.Context, err error, message string) {
	l.Fatal(GetRequestID(ctx), err, ErrMetadata(err), message)
}
