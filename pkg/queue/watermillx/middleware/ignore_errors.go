package middleware

import (
	"strings"

	"github.com/ThreeDotsLabs/watermill/message"
)

var ignoreErrors = []string{
	"uuid.Parse",
	"Could not decode AMQP frame",
}

// NewIgnoreErrors creates a new IgnoreErrors middleware.
func NewIgnoreErrors() message.HandlerMiddleware {
	errsMap := make(map[string]bool, len(ignoreErrors))

	for _, err := range ignoreErrors {
		errsMap[err] = true
	}

	unit := IgnoreErrors{errsMap}
	return unit.Middleware
}

// IgnoreErrors provides a middleware that makes the handler ignore some explicitly whitelisted errors.
type IgnoreErrors struct {
	list map[string]bool
}

func (i IgnoreErrors) Middleware(h message.HandlerFunc) message.HandlerFunc {
	return func(msg *message.Message) ([]*message.Message, error) {
		events, err := h(msg)
		if err == nil {
			return events, nil
		}

		for k := range i.list {
			if strings.Contains(err.Error(), k) {
				return events, nil
			}
		}

		return events, err

	}
}
