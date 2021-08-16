package nsqx

import (
	"context"
	"time"
)

//go:generate mockgen -destination=nsqx_mock.go -package=nsqx -source=nsqx.go

// PublisherItf is publisher interface to publish nsq message.
type PublisherItf interface {
	// Publish sends data to nsq.
	// Data should be json bytes but the struct or map.
	Publish(ctx context.Context, topic string, data interface{}) error

	// DeferredPublish sends data to nsq after certain delay.
	DeferredPublish(ctx context.Context, topic string, delay time.Duration, data interface{}) error
}
