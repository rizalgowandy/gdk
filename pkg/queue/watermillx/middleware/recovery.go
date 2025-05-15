package middleware

import (
	"fmt"
	"runtime/debug"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

type RecoveryPanicError struct {
	V          any
	Stacktrace string
}

func (p RecoveryPanicError) Error() string {
	return fmt.Sprintf("panic occurred: %#v, stacktrace: \n%s", p.V, p.Stacktrace)
}

// NewRecovery recovers from any panic in the handler and appends RecoveryPanicError with the stacktrace
// to any error returned from the handler.
func NewRecovery(h message.HandlerFunc) message.HandlerFunc {
	return func(event *message.Message) (events []*message.Message, err error) {
		defer func() {
			if r := recover(); r != nil {
				panicErr := errors.WithStack(
					RecoveryPanicError{
						V:          r,
						Stacktrace: string(debug.Stack()),
					},
				)
				err = multierror.Append(err, panicErr)
			}
		}()

		return h(event)
	}
}
