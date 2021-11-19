package middleware

import (
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

func NewRetry(logger watermill.LoggerAdapter) message.HandlerMiddleware {
	unit := middleware.Retry{
		MaxRetries:      5,
		InitialInterval: time.Millisecond * 100,
		Logger:          logger,
	}
	return unit.Middleware
}
