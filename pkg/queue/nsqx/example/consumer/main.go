package main

import (
	"context"

	"github.com/nsqio/go-nsq"
	"github.com/peractio/gdk/pkg/logx"
	"github.com/peractio/gdk/pkg/queue/nsqx"
	"github.com/peractio/gdk/pkg/queue/nsqx/interceptor"
)

type SendEmail struct{}

func (s SendEmail) Handle(ctx context.Context, message *nsq.Message) error {
	logx.INF(ctx, nil, "send email is running")
	return nil
}

func main() {
	// Create controller.
	ctrl := nsqx.NewConsumerController(
		nsqx.ConsumerChain(
			interceptor.RequestID,
			interceptor.Recover(),
			interceptor.Logger,
		),
	)

	// Add consumers.
	ctx := context.Background()
	if err := ctrl.AddConsumers(GetConsumers()); err != nil {
		logx.FTL(ctx, err, "add consumers failure")
	}

	ctrl.Serve()
}

func GetConsumers() []nsqx.ConsumerParam {
	return []nsqx.ConsumerParam{
		{
			Topic:   "new_user_registered",
			Channel: "send_email",
			Config: &nsqx.ConsumerConfiguration{
				NSQ:           nsq.NewConfig(),
				LookupAddress: []string{"127.0.0.1:4161"},
				Concurrency:   100,
				MaxInFlight:   100,
				MaxAttempts:   5,
			},
			Consumer: SendEmail{},
		},
	}
}
