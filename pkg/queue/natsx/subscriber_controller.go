package natsx

import (
	"os"
	"os/signal"
	"sync"

	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
)

type SubscriberController struct {
	conn        *nats.Conn
	subscribers []*nats.Subscription
	interceptor SubscriberInterceptor
}

type SubscriberParam struct {
	Subject    string
	Queue      string
	Subscriber Subscriber
}

func NewSubscriberController(conn *nats.Conn, interceptors ...SubscriberInterceptor) *SubscriberController {
	return &SubscriberController{
		conn:        conn,
		subscribers: nil,
		interceptor: SubscriberChain(interceptors...),
	}
}

func (c *SubscriberController) AddQueueSubscriber(params []SubscriberParam) error {
	for _, param := range params {
		if param.Subscriber == nil {
			return ErrInvalidSubscriber
		}

		subscriber, err := c.conn.QueueSubscribe(
			param.Subject,
			param.Queue,
			NewSubscriber(c, param.Subject, param.Queue, param.Subscriber).HandleMessage,
		)
		if err != nil {
			return errors.Wrap(err, "conn.QueueSubscribe")
		}

		c.subscribers = append(c.subscribers, subscriber)
	}

	return nil
}

func (c *SubscriberController) Serve() {
	finish := make(chan struct{})
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit

		var wg sync.WaitGroup
		for _, v := range c.subscribers {
			wg.Add(1)
			con := v

			go func() {
				defer wg.Done()
				_ = con.Unsubscribe()
			}()
		}
		wg.Wait()
		close(finish)
	}()
	<-finish
}
