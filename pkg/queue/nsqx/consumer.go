package nsqx

import (
	"context"
	"fmt"

	"github.com/nsqio/go-nsq"
)

type Consumer struct {
	Topic   string
	Channel string
	Method  string

	config *ConsumerConfiguration
	ctrl   *ConsumerController
	inner  ConsumerItf
}

func (c *Consumer) HandleMessage(message *nsq.Message) error {
	ctx := context.Background()

	err := c.ctrl.interceptor(ctx, c, func(ctx context.Context, consumer *Consumer) error {
		return consumer.inner.Handle(ctx, message)
	})

	return err
}

func NewConsumer(
	ctrl *ConsumerController,
	topic, channel string,
	config *ConsumerConfiguration,
	consumer ConsumerItf,
) *Consumer {
	return &Consumer{
		Topic:   topic,
		Channel: channel,
		Method:  fmt.Sprintf("%s-%s", topic, channel),
		config:  config,
		ctrl:    ctrl,
		inner:   consumer,
	}
}

// FuncConsumer is a type to allow callers to wrap a raw func.
type FuncConsumer func(ctx context.Context, message *nsq.Message) error

func (r FuncConsumer) Handle(ctx context.Context, message *nsq.Message) error {
	return r(ctx, message)
}
