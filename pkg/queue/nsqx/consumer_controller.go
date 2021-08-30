package nsqx

import (
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/rizalgowandy/gdk/pkg/errorx/v2"
)

const TimeoutDuration = time.Second * 60

type ConsumerController struct {
	consumers   []*nsq.Consumer
	interceptor ConsumerInterceptor
}

type ConsumerParam struct {
	Topic    string
	Channel  string
	Config   *ConsumerConfiguration
	Consumer ConsumerItf
}

func NewConsumerController(interceptors ...ConsumerInterceptor) *ConsumerController {
	return &ConsumerController{
		consumers:   nil,
		interceptor: ConsumerChain(interceptors...),
	}
}

func (c *ConsumerController) AddConsumers(params []ConsumerParam) error {
	const op errorx.Op = "nsqx/ConsumerController.AddConsumers"

	for _, param := range params {
		if param.Config == nil {
			param.Config = &ConsumerConfiguration{}
		}
		if param.Consumer == nil {
			return errorx.E("invalid consumer", op)
		}

		if err := param.Config.Validate(); err != nil {
			return errorx.E(err, op)
		}

		consumer, err := nsq.NewConsumer(param.Topic, param.Channel, param.Config.NSQ)
		if err != nil {
			return errorx.E(err, op)
		}

		consumer.AddConcurrentHandlers(
			NewConsumer(
				c,
				param.Topic,
				param.Channel,
				param.Config,
				param.Consumer,
			),
			param.Config.Concurrency,
		)

		err = consumer.ConnectToNSQLookupds(param.Config.LookupAddress)
		if err != nil {
			return errorx.E(err, op)
		}

		c.consumers = append(c.consumers, consumer)
	}

	return nil
}

func (c *ConsumerController) Serve() {
	finish := make(chan struct{})
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit

		var wg sync.WaitGroup
		for _, v := range c.consumers {
			wg.Add(1)
			con := v

			go func() {
				defer wg.Done()
				con.Stop()

				select {
				case <-con.StopChan:
				case <-time.After(TimeoutDuration):
				}
			}()
		}
		wg.Wait()
		close(finish)
	}()
	<-finish
}
