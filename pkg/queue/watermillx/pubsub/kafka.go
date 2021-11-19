package pubsub

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/segmentio/ksuid"
)

func NewKafkaSubscriber(
	logger watermill.LoggerAdapter,
	addresses []string,
) (*kafka.Subscriber, error) {
	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:       addresses,
			Unmarshaler:   kafka.DefaultMarshaler{},
			ConsumerGroup: ksuid.New().String(),
		},
		logger,
	)
	if err != nil {
		return nil, err
	}

	return subscriber, nil
}

func NewKafkaPublisher(
	logger watermill.LoggerAdapter,
	addresses []string,
) (*kafka.Publisher, error) {
	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   addresses,
			Marshaler: kafka.DefaultMarshaler{},
		},
		logger,
	)
	if err != nil {
		return nil, err
	}

	return publisher, nil
}
