package nsqx

import (
	"github.com/nsqio/go-nsq"
	"github.com/peractio/gdk/pkg/errorx/v2"
)

type PublisherConfiguration struct {
	NSQ           *nsq.Config
	DaemonAddress string
	MaxAttempt    int
}

func (p *PublisherConfiguration) Validate() error {
	if p.NSQ == nil {
		p.NSQ = nsq.NewConfig()
	}
	if p.DaemonAddress == "" {
		return errorx.New("missing daemon address")
	}
	if p.MaxAttempt <= 1 {
		p.MaxAttempt = 1
	}
	return nil
}

type ConsumerConfiguration struct {
	NSQ           *nsq.Config
	LookupAddress []string
	Concurrency   int
	MaxInFlight   int
	MaxAttempts   uint16
}

func (c *ConsumerConfiguration) Validate() error {
	if c.NSQ == nil {
		c.NSQ = nsq.NewConfig()
	}
	if c.Concurrency <= 0 {
		c.Concurrency = 1
	}
	if c.MaxInFlight > 0 {
		c.NSQ.MaxInFlight = c.MaxInFlight
	}
	if c.MaxAttempts > 0 {
		c.NSQ.MaxAttempts = c.MaxAttempts
	}

	return nil
}
