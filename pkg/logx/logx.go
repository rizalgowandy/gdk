package logx

import (
	"context"

	"github.com/peractio/gdk/pkg/syncx"
)

// KV is wrapper for map[string]interface{}.
type KV map[string]interface{}

type Config struct {
	Debug    bool
	AppName  string
	Filename string
}

var (
	once    syncx.Once
	onceRes Logger
	onceErr error

	// DefaultConfig is a configuration to create:
	// - Non JSON structured.
	// - No app name.
	// - Output to os.Stderr.
	DefaultConfig = Config{
		Debug:    true,
		AppName:  "",
		Filename: "",
	}
)

func New(configs ...Config) error {
	once.Do(func() {
		var config Config
		if len(configs) == 0 {
			config = DefaultConfig
		} else {
			config = configs[0]
		}

		logger, err := NewZerolog(config)
		if err != nil {
			onceErr = err
			return
		}
		onceRes = logger
	})
	return onceErr
}

func TRC(ctx context.Context, metadata interface{}, message string) {
	if createErr := New(); createErr != nil {
		return
	}
	onceRes.Trace(GetRequestID(ctx), Metadata(metadata), message)
}

func DBG(ctx context.Context, metadata interface{}, message string) {
	if createErr := New(); createErr != nil {
		return
	}
	onceRes.Debug(GetRequestID(ctx), Metadata(metadata), message)
}

func INF(ctx context.Context, metadata interface{}, message string) {
	if createErr := New(); createErr != nil {
		return
	}
	onceRes.Info(GetRequestID(ctx), Metadata(metadata), message)
}

func WRN(ctx context.Context, err error, message string) {
	if createErr := New(); createErr != nil {
		return
	}
	onceRes.Warn(GetRequestID(ctx), err, ErrMetadata(err), message)
}

func ERR(ctx context.Context, err error, message string) {
	if createErr := New(); createErr != nil {
		return
	}
	onceRes.Error(GetRequestID(ctx), err, ErrMetadata(err), message)
}

func FTL(ctx context.Context, err error, message string) {
	if createErr := New(); createErr != nil {
		return
	}
	onceRes.Fatal(GetRequestID(ctx), err, ErrMetadata(err), message)
}
