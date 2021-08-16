package nsqx

import (
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/peractio/gdk/pkg/errorx/v2"
)

const TimeoutDuration = time.Second * 60

type ConsumerCommander struct {
	cons []*nsq.Consumer
}

// NewConsumerCommander creates new consumer commander.
func NewConsumerCommander() *ConsumerCommander {
	return &ConsumerCommander{}
}

func (c *ConsumerCommander) AddHandler(
	topic, channel string,
	cfg *ConsumerConfiguration,
	handler nsq.Handler,
) error {
	const op errorx.Op = "nsqx/ConsumerCommander.AddHandler"

	if err := cfg.Validate(); err != nil {
		return errorx.E(err, op)
	}

	con, err := nsq.NewConsumer(topic, channel, cfg.NSQ)
	if err != nil {
		return errorx.E(err, op)
	}

	con.AddConcurrentHandlers(handler, cfg.Concurrency)

	err = con.ConnectToNSQLookupds(cfg.LookupAddress)
	if err != nil {
		return errorx.E(err, op)
	}

	c.cons = append(c.cons, con)
	return nil
}

func (c *ConsumerCommander) Serve() <-chan struct{} {
	finish := make(chan struct{})
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit

		var wg sync.WaitGroup
		for _, con := range c.cons {
			wg.Add(1)

			con := con
			go func() {
				defer wg.Done()
				(*con).Stop()

				select {
				case <-con.StopChan:
				case <-time.After(TimeoutDuration):
				}
			}()
		}
		wg.Wait()
		close(finish)
	}()
	return finish
}
