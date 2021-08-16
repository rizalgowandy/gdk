package interceptor

import (
	"context"
	"testing"

	"github.com/peractio/gdk/pkg/errorx/v2"
	"github.com/peractio/gdk/pkg/logx"
	"github.com/peractio/gdk/pkg/queue/nsqx"
)

func TestLogger(t *testing.T) {
	type args struct {
		ctx      context.Context
		consumer *nsqx.Consumer
		handler  nsqx.ConsumerHandler
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Error",
			args: args{
				ctx:      context.Background(),
				consumer: &nsqx.Consumer{},
				handler: func(ctx context.Context, consumer *nsqx.Consumer) error {
					return errorx.New("error")
				},
			},
			wantErr: true,
		},
		{
			name: "Success",
			args: args{
				ctx:      context.Background(),
				consumer: &nsqx.Consumer{},
				handler: func(ctx context.Context, consumer *nsqx.Consumer) error {
					return nil
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _ = logx.New(&logx.Config{
				Debug:    true,
				AppName:  "unit_test",
				Filename: "",
			})

			if err := Logger(tt.args.ctx, tt.args.consumer, tt.args.handler); (err != nil) != tt.wantErr {
				t.Errorf("Logger() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
