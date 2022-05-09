package nsqx

import (
	"context"
	"sync"
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/rizalgowandy/gdk/pkg/errorx/v2"
	"github.com/rizalgowandy/gdk/pkg/jsonx"
	"github.com/rizalgowandy/gdk/pkg/syncx"
)

var (
	onceNewProducer    syncx.Once
	onceNewProducerRes *Producer
	onceNewProducerErr error
)

type Producer struct {
	config *ProducerConfiguration
	client *nsq.Producer
	mux    sync.Mutex
}

// NewProducer creates a client to publish message to nsq.
func NewProducer(config *ProducerConfiguration) (*Producer, error) {
	onceNewProducer.Do(func() {
		if err := config.Validate(); err != nil {
			onceNewProducerErr = errorx.E(err)
			return
		}

		client, err := nsq.NewProducer(config.DaemonAddress, config.NSQ)
		if err != nil {
			onceNewProducerErr = errorx.E(err)
			return
		}

		onceNewProducerRes = &Producer{
			config: config,
			client: client,
		}
	})

	return onceNewProducerRes, onceNewProducerErr
}

// Publish sends data to nsq.
func (p *Producer) Publish(_ context.Context, topic string, data interface{}) error {
	if topic == "" {
		return errorx.E("topic cannot be empty")
	}

	var err error
	dataByte, ok := (data).([]byte)
	if !ok {
		dataByte, err = jsonx.Marshal(data)
		if err != nil {
			return errorx.E(err)
		}
	}

	err = p.publish(topic, dataByte)
	if err != nil {
		return errorx.E(err)
	}

	return nil
}

// DeferredPublish sends data to nsq after certain delay.
func (p *Producer) DeferredPublish(
	_ context.Context,
	topic string,
	delay time.Duration,
	data interface{},
) error {
	if topic == "" {
		return errorx.E("topic cannot be empty")
	}

	var err error
	dataByte, ok := (data).([]byte)
	if !ok {
		dataByte, err = jsonx.Marshal(data)
		if err != nil {
			return errorx.E(err)
		}
	}

	err = p.deferredPublish(topic, delay, dataByte)
	if err != nil {
		return errorx.E(err)
	}

	return nil
}

func (p *Producer) publish(topic string, data []byte) error {
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

func (p *Producer) deferredPublish(topic string, delay time.Duration, data []byte) error {
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
