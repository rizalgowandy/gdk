package natsx

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
)

type SubscriberImpl struct {
	Subject string
	Queue   string
	Method  string

	ctrl  *SubscriberController
	inner Subscriber
}

//nolint:wrapcheck
func (c *SubscriberImpl) HandleMessage(message *nats.Msg) {
	ctx := context.Background()
	_ = c.ctrl.interceptor(ctx, c, func(ctx context.Context, subscriber *SubscriberImpl) error {
		return subscriber.inner.Handle(ctx, message)
	})
}

func NewSubscriber(
	ctrl *SubscriberController,
	subj, queue string,
	subscriber Subscriber,
) *SubscriberImpl {
	return &SubscriberImpl{
		Subject: subj,
		Queue:   queue,
		Method:  fmt.Sprintf("%s-%s", subj, queue),
		ctrl:    ctrl,
		inner:   subscriber,
	}
}
