package pubsub

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
)

func NewAMQPSubscriber(
	logger watermill.LoggerAdapter,
	address string,
) (*amqp.Subscriber, error) {
	subscriber, err := amqp.NewSubscriber(
		amqp.NewDurableQueueConfig(address),
		logger,
	)
	if err != nil {
		return nil, err
	}

	return subscriber, nil
}

func NewAMQPPublisher(
	logger watermill.LoggerAdapter,
	address string,
) (*amqp.Publisher, error) {
	publisher, err := amqp.NewPublisher(
		amqp.NewDurableQueueConfig(address),
		logger,
	)
	if err != nil {
		return nil, err
	}

	return publisher, nil
}
