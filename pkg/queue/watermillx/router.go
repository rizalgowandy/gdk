package watermillx

import (
	"context"
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/rizalgowandy/gdk/pkg/queue/watermillx/middleware"
)

type Handler struct {
	Topic      string
	Channel    string
	Subscriber message.Subscriber
	Exec       message.NoPublishHandlerFunc
}

func NewRouter(
	_ context.Context,
	logger watermill.LoggerAdapter,
	handlers []Handler,
) (*message.Router, error) {
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return nil, err
	}

	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(middleware.New(logger)...)

	for _, v := range handlers {
		router.AddNoPublisherHandler(
			fmt.Sprintf("%s_%s", v.Topic, v.Channel),
			v.Topic,
			v.Subscriber,
			v.Exec,
		)
	}

	return router, nil
}
