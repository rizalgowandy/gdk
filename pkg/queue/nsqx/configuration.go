package nsqx

import (
	"github.com/nsqio/go-nsq"
	"github.com/peractio/gdk/pkg/errorx/v2"
)

type Configuration struct {
	DaemonAddress string
	MaxRetry      int
	NSQ           *nsq.Config
}

func (c *Configuration) Validate() error {
	if c.DaemonAddress == "" {
		return errorx.New("missing daemon address")
	}
	if c.MaxRetry <= 1 {
		c.MaxRetry = 1
	}
	if c.NSQ == nil {
		c.NSQ = nsq.NewConfig()
	}
	return nil
}
