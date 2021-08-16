package nsqx

import (
	"github.com/peractio/gdk/pkg/errorx/v2"
)

type ConsumerController struct {
	Commander *ConsumerCommander

	// Interceptor holds middleware that will be executed before current consumer operation.
	Interceptor ConsumerInterceptor
}

type ConsumerParam struct {
	Topic    string
	Channel  string
	Config   *ConsumerConfiguration
	Consumer ConsumerItf
}

func NewConsumerController(interceptors ...ConsumerInterceptor) *ConsumerController {
	return &ConsumerController{
		Commander:   NewConsumerCommander(),
		Interceptor: ConsumerChain(interceptors...),
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

		if err := c.Commander.AddHandler(
			param.Topic,
			param.Channel,
			param.Config,
			NewConsumer(
				c,
				param.Topic,
				param.Channel,
				param.Config,
				param.Consumer,
			),
		); err != nil {
			return errorx.E(err, op)
		}
	}

	return nil
}

func (c *ConsumerController) Serve() {
	<-c.Commander.Serve()
}
