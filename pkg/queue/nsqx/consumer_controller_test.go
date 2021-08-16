package nsqx

import (
	"context"
	"testing"

	"github.com/nsqio/go-nsq"
	"github.com/stretchr/testify/assert"
)

func TestNewConsumerController(t *testing.T) {
	type args struct {
		interceptors []ConsumerInterceptor
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewConsumerController(tt.args.interceptors...)
			assert.NotNil(t, got)
		})
	}
}

func TestConsumerController_AddConsumers(t *testing.T) {
	type fields struct {
		Interceptor ConsumerInterceptor
	}
	type args struct {
		params []ConsumerParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Error missing consumer",
			fields: fields{
				Interceptor: nil,
			},
			args: args{
				params: []ConsumerParam{
					{
						Topic:    "Topic",
						Channel:  "Channel",
						Config:   nil,
						Consumer: nil,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Success",
			fields: fields{
				Interceptor: nil,
			},
			args: args{
				params: []ConsumerParam{
					{
						Topic:   "Topic",
						Channel: "Channel",
						Config:  nil,
						Consumer: FuncConsumer(
							func(ctx context.Context, message *nsq.Message) error {
								return nil
							},
						),
					},
					{
						Topic:   "Topic",
						Channel: "Channel",
						Config:  nil,
						Consumer: FuncConsumer(
							func(ctx context.Context, message *nsq.Message) error {
								return nil
							},
						),
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ConsumerController{
				Interceptor: tt.fields.Interceptor,
			}
			if err := c.AddConsumers(tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("AddConsumers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
