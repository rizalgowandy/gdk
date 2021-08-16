package nsqx

import (
	"context"
	"sync"
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/peractio/gdk/pkg/errorx/v2"
	"github.com/peractio/gdk/pkg/jsonx"
	"github.com/peractio/gdk/pkg/syncx"
)

var (
	onceNewPublisher    syncx.Once
	onceNewPublisherRes *Publisher
	onceNewPublisherErr error
)

type Publisher struct {
	config *PublisherConfiguration
	client *nsq.Producer
	mux    sync.Mutex
}

// NewPublisher creates a client to publish message to nsq.
func NewPublisher(config *PublisherConfiguration) (*Publisher, error) {
	onceNewPublisher.Do(func() {
		const op errorx.Op = "nsqx.NewPublisher"

		if err := config.Validate(); err != nil {
			onceNewPublisherErr = errorx.E(err, op)
			return
		}

		client, err := nsq.NewProducer(config.DaemonAddress, config.NSQ)
		if err != nil {
			onceNewPublisherErr = errorx.E(err, op)
			return
		}

		onceNewPublisherRes = &Publisher{
			config: config,
			client: client,
		}
	})

	return onceNewPublisherRes, onceNewPublisherErr
}

// Publish sends data to nsq.
func (p *Publisher) Publish(_ context.Context, topic string, data interface{}) error {
	const op errorx.Op = "nsqx/Publisher.Publish"

	if topic == "" {
		return errorx.E("topic cannot be empty", op)
	}

	var err error
	dataByte, ok := (data).([]byte)
	if !ok {
		dataByte, err = jsonx.Marshal(data)
		if err != nil {
			return errorx.E(err, op)
		}
	}

	err = p.publish(topic, dataByte)
	if err != nil {
		return errorx.E(err, op)
	}

	return nil
}

// DeferredPublish sends data to nsq after certain delay.
func (p *Publisher) DeferredPublish(
	_ context.Context,
	topic string,
	delay time.Duration,
	data interface{},
) error {
	const op errorx.Op = "nsqx/Publisher.DeferredPublish"

	if topic == "" {
		return errorx.E("topic cannot be empty", op)
	}

	var err error
	dataByte, ok := (data).([]byte)
	if !ok {
		dataByte, err = jsonx.Marshal(data)
		if err != nil {
			return errorx.E(err, op)
		}
	}

	err = p.deferredPublish(topic, delay, dataByte)
	if err != nil {
		return errorx.E(err, op)
	}

	return nil
}

func (p *Publisher) publish(topic string, data []byte) error {
	var err error

	for i := 1; i <= p.config.MaxAttempt; i++ {
		err = func() error {
			p.mux.Lock()
			err = p.client.Publish(topic, data)
			p.mux.Unlock()
			return err
		}()

		// Success.
		if err == nil {
			return nil
		}

		// Max retry achieved.
		if i == p.config.MaxAttempt {
			break
		}

		func() {
			p.mux.Lock()
			defer p.mux.Unlock()

			prod, newErr := nsq.NewProducer(p.config.DaemonAddress, p.config.NSQ)
			if newErr == nil {
				p.client.Stop()
				p.client = prod
			}
		}()
	}

	return err
}

func (p *Publisher) deferredPublish(topic string, delay time.Duration, data []byte) error {
	var err error

	for i := 1; i <= p.config.MaxAttempt; i++ {
		err = func() error {
			p.mux.Lock()
			err = p.client.DeferredPublish(topic, delay, data)
			p.mux.Unlock()
			return err
		}()

		// Success.
		if err == nil {
			return nil
		}

		// Max retry achieved.
		if i == p.config.MaxAttempt {
			break
		}

		func() {
			p.mux.Lock()
			defer p.mux.Unlock()

			prod, newErr := nsq.NewProducer(p.config.DaemonAddress, p.config.NSQ)
			if newErr == nil {
				p.client.Stop()
				p.client = prod
			}
		}()
	}

	return err
}
