package nsqx

import (
	"context"
	"testing"

	"github.com/nsqio/go-nsq"
	"github.com/stretchr/testify/assert"
)

func TestConsumer_HandleMessage(t *testing.T) {
	type fields struct {
		ctrl    *ConsumerController
		topic   string
		channel string
		config  *ConsumerConfiguration
		inner   ConsumerItf
	}
	type args struct {
		message *nsq.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				ctrl: &ConsumerController{
					interceptor: ConsumerChain(),
				},
				topic:   "Topic",
				channel: "Channel",
				config:  &ConsumerConfiguration{},
				inner: FuncConsumer(func(ctx context.Context, message *nsq.Message) error {
					return nil
				}),
			},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Consumer{
				ctrl:    tt.fields.ctrl,
				Topic:   tt.fields.topic,
				Channel: tt.fields.channel,
				config:  tt.fields.config,
				inner:   tt.fields.inner,
			}
			if err := c.HandleMessage(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("HandleMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncConsumer_Handle(t *testing.T) {
	type args struct {
		ctx     context.Context
		message *nsq.Message
	}
	tests := []struct {
		name    string
		r       FuncConsumer
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			r: FuncConsumer(func(ctx context.Context, message *nsq.Message) error {
				return nil
			}),
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Handle(tt.args.ctx, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("Handle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewConsumer(t *testing.T) {
	type args struct {
		ctrl     *ConsumerController
		topic    string
		channel  string
		config   *ConsumerConfiguration
		consumer ConsumerItf
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
			got := NewConsumer(
				tt.args.ctrl,
				tt.args.topic,
				tt.args.channel,
				tt.args.config,
				tt.args.consumer,
			)
			assert.NotNil(t, got)
		})
	}
}
