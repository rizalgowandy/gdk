package watermillx

import (
	"context"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/rizalgowandy/gdk/pkg/errorx/v2"
	"github.com/rizalgowandy/gdk/pkg/logx"
)

func NewLogger() (*Logger, error) {
	return &Logger{}, nil
}

type Logger struct {
	fields watermill.LogFields
}

func (l *Logger) Error(msg string, err error, fields watermill.LogFields) {
	logx.ERR(context.Background(), l.ErrMetadata(err, fields), msg)
}

func (l *Logger) Info(msg string, fields watermill.LogFields) {
	logx.INF(context.Background(), fields, msg)
}

func (l *Logger) Debug(msg string, fields watermill.LogFields) {
	logx.DBG(context.Background(), fields, msg)
}

func (l *Logger) Trace(msg string, fields watermill.LogFields) {
	logx.TRC(context.Background(), fields, msg)
}

func (l *Logger) With(fields watermill.LogFields) watermill.LoggerAdapter {
	l.fields = fields
	return l
}

func (l *Logger) ErrMetadata(err error, fields watermill.LogFields) error {
	md := make(logx.KV)
	for k, v := range fields {
		md[k] = v
	}
	return errorx.E(err, md)
}
