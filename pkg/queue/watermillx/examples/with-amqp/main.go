package main

import (
	"log"

	"github.com/rizalgowandy/gdk/pkg/logx"
	"github.com/rizalgowandy/gdk/pkg/queue/watermillx"
	"github.com/rizalgowandy/gdk/pkg/queue/watermillx/examples/with-amqp/job"
	"github.com/rizalgowandy/gdk/pkg/queue/watermillx/examples/with-amqp/topic"
	"github.com/rizalgowandy/gdk/pkg/queue/watermillx/pubsub"
)

const consumerName = "default"

func main() {
	ctx := logx.NewContext()

	logger, err := watermillx.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	address := "amqp://guest:guest@localhost:5672/"

	// Create publisher.
	publisher, err := pubsub.NewAMQPPublisher(logger, address)
	if err != nil {
		log.Fatal(err)
	}

	// Create subscriber.
	subscriber, err := pubsub.NewAMQPSubscriber(logger, address)
	if err != nil {
		log.Fatal(err)
	}

	// Create a list of workers.
	handlers := []watermillx.Handler{
		{
			Topic:      topic.A,
			Channel:    consumerName,
			Subscriber: subscriber,
			Exec:       job.NewA(ctx, publisher),
		},
		{
			Topic:      topic.B,
			Channel:    consumerName,
			Subscriber: subscriber,
			Exec:       job.NewB(ctx),
		},
	}

	// Create router with middleware.
	router, err := watermillx.NewRouter(ctx, logger, handlers)
	if err != nil {
		log.Fatal(err)
	}

	// Run subscriber for real.
	// This call is blocking while the router is running.
	err = router.Run(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
