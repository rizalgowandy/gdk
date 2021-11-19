package main

import (
	"log"

	"github.com/rizalgowandy/gdk/pkg/logx"
	"github.com/rizalgowandy/gdk/pkg/queue/watermillx"
	"github.com/rizalgowandy/gdk/pkg/queue/watermillx/examples/with-kafka/job"
	"github.com/rizalgowandy/gdk/pkg/queue/watermillx/examples/with-kafka/topic"
	"github.com/rizalgowandy/gdk/pkg/queue/watermillx/pubsub"
)

const consumerName = "default"

func main() {
	ctx := logx.NewContext()

	logger, err := watermillx.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	address := []string{"kafka:9092"}

	// Create publisher.
	publisher, err := pubsub.NewKafkaPublisher(logger, address)
	if err != nil {
		log.Fatal(err)
	}

	// Create subscriber.
	subscriber, err := pubsub.NewKafkaSubscriber(logger, address)
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
