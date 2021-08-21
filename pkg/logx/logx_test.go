package logx

import (
	"context"
	"testing"

	"github.com/peractio/gdk/pkg/errorx/v2"
)

func TestTRC(t *testing.T) {
	type args struct {
		ctx      context.Context
		metadata interface{}
		message  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				ctx: NewContext(),
				metadata: map[string]string{
					"k": "v",
				},
				message: "trace log",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TRC(tt.args.ctx, tt.args.metadata, tt.args.message)
		})
	}
}

func TestDBG(t *testing.T) {
	type args struct {
		ctx      context.Context
		metadata interface{}
		message  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				ctx: NewContext(),
				metadata: map[string]string{
					"k": "v",
				},
				message: "trace log",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DBG(tt.args.ctx, tt.args.metadata, tt.args.message)
		})
	}
}

func TestINF(t *testing.T) {
	type args struct {
		ctx      context.Context
		metadata interface{}
		message  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				ctx: NewContext(),
				metadata: map[string]string{
					"k": "v",
				},
				message: "trace log",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			INF(tt.args.ctx, tt.args.metadata, tt.args.message)
		})
	}
}

func TestWRN(t *testing.T) {
	type args struct {
		ctx     context.Context
		err     error
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				ctx: NewContext(),
				err: errorx.E("random error", errorx.CodeInternal, errorx.Fields{
					"k": "v",
				}, errorx.MetricStatusExpectedErr, errorx.Message("human message")),
				message: "error log",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WRN(tt.args.ctx, tt.args.err, tt.args.message)
		})
	}
}

func TestERR(t *testing.T) {
	type args struct {
		ctx     context.Context
		err     error
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{
				ctx:     NewContext(),
				err:     errorx.E("random error"),
				message: "error log",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ERR(tt.args.ctx, tt.args.err, tt.args.message)
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Without configuration",
			args:    args{},
			wantErr: false,
		},
		{
			name: "Success",
			args: args{
				config: Config{
					Debug:    true,
					AppName:  "unit_test",
					Filename: "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := New(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
