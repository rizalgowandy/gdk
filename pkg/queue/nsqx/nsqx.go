package nsqx

import (
	"context"
	"time"

	"github.com/nsqio/go-nsq"
)

//go:generate mockgen -destination=nsqx_mock.go -package=nsqx -source=nsqx.go

// ProducerItf is producer interface to publish nsq message.
type ProducerItf interface {
	// Publish sends data to nsq.
	// Data should be json bytes but the struct or map.
	Publish(ctx context.Context, topic string, data any) error

	// DeferredPublish sends data to nsq after certain delay.
	DeferredPublish(ctx context.Context, topic string, delay time.Duration, data any) error
}

// ConsumerItf is consumer interface to consume nsq message.
type ConsumerItf interface {
	Handle(ctx context.Context, message *nsq.Message) error
}
