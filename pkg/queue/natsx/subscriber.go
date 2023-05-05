package natsx

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
)

type Subscriber struct {
	Subject string
	Queue   string
	Method  string

	ctrl  *SubscriberController
	inner SubscriberItf
}

func (c *Subscriber) HandleMessage(message *nats.Msg) {
	ctx := context.Background()
	_ = c.ctrl.interceptor(ctx, c, func(ctx context.Context, subscriber *Subscriber) error {
		return subscriber.inner.Handle(ctx, message)
	})
}

func NewSubscriber(
	ctrl *SubscriberController,
	subj, queue string,
	subscriber SubscriberItf,
) *Subscriber {
	return &Subscriber{
		Subject: subj,
		Queue:   queue,
		Method:  fmt.Sprintf("%s-%s", subj, queue),
		ctrl:    ctrl,
		inner:   subscriber,
	}
}
