package middleware

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

// New returns the middleware interceptor.
// Remember that the order of the interceptor is important.
// The first one is gonna be executed first.
func New(
	logger watermill.LoggerAdapter,
) []message.HandlerMiddleware {
	return []message.HandlerMiddleware{
		NewRecovery,
		middleware.CorrelationID,
		NewMetadata,
		NewRetry(logger),
		NewIgnoreErrors(),
	}
}
