package logx

import (
	"testing"

	"github.com/rizalgowandy/gdk/pkg/errorx/v2"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestNewZerolog(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success with debug",
			args: args{
				config: Config{
					Debug:    true,
					AppName:  "gdk",
					Filename: "",
				},
			},
			wantErr: false,
		},
		{
			name: "Success",
			args: args{
				config: Config{
					Debug:    false,
					AppName:  "gdk",
					Filename: "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewZerolog(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewZerolog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				assert.NotNil(t, got)
			} else {
				assert.Nil(t, got)
			}
		})
	}
}

func TestZeroLogger_Trace(t *testing.T) {
	type fields struct {
		client zerolog.Logger
	}
	type args struct {
		requestID string
		fields    map[string]any
		message   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Success",
			fields: fields{
				client: zerolog.Logger{},
			},
			args: args{
				requestID: "",
				fields:    nil,
				message:   "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &ZeroLogger{
				client: tt.fields.client,
			}
			z.Trace(tt.args.requestID, tt.args.fields, tt.args.message)
		})
	}
}

func TestZeroLogger_Debug(t *testing.T) {
	type fields struct {
		client zerolog.Logger
	}
	type args struct {
		requestID string
		fields    map[string]any
		message   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Success",
			fields: fields{
				client: zerolog.Logger{},
			},
			args: args{
				requestID: "",
				fields:    nil,
				message:   "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &ZeroLogger{
				client: tt.fields.client,
			}
			z.Debug(tt.args.requestID, tt.args.fields, tt.args.message)
		})
	}
}

func TestZeroLogger_Info(t *testing.T) {
	type fields struct {
		client zerolog.Logger
	}
	type args struct {
		requestID string
		fields    map[string]any
		message   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Success",
			fields: fields{
				client: zerolog.Logger{},
			},
			args: args{
				requestID: "",
				fields:    nil,
				message:   "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &ZeroLogger{
				client: tt.fields.client,
			}
			z.Info(tt.args.requestID, tt.args.fields, tt.args.message)
		})
	}
}

func TestZeroLogger_Warn(t *testing.T) {
	type fields struct {
		client zerolog.Logger
	}
	type args struct {
		requestID string
		err       error
		fields    map[string]any
		message   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Success",
			fields: fields{
				client: zerolog.Logger{},
			},
			args: args{
				requestID: "",
				err:       errorx.E("abc"),
				fields:    nil,
				message:   "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &ZeroLogger{
				client: tt.fields.client,
			}
			z.Warn(tt.args.requestID, tt.args.err, tt.args.fields, tt.args.message)
		})
	}
}

func TestZeroLogger_Error(t *testing.T) {
	type fields struct {
		client zerolog.Logger
	}
	type args struct {
		requestID string
		err       error
		fields    map[string]any
		message   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Success",
			fields: fields{
				client: zerolog.Logger{},
			},
			args: args{
				requestID: "",
				err:       errorx.E("abc"),
				fields:    nil,
				message:   "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &ZeroLogger{
				client: tt.fields.client,
			}
			z.Error(tt.args.requestID, tt.args.err, tt.args.fields, tt.args.message)
		})
	}
}
