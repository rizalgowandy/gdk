package natsx

import (
	"context"

	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
)

var (
	ErrInvalidSubscriber = errors.New("invalid subscriber")
)

// Subscriber is subscriber interface to consume nats message.
type Subscriber interface {
	Handle(ctx context.Context, message *nats.Msg) error
}
