package natsx

import (
	"context"

	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
)

//go:generate mockgen -destination=natsx_mock.go -package=natsx -source=natsx.go

var (
	ErrInvalidSubscriber = errors.New("invalid subscriber")
)

// SubscriberItf is subscriber interface to consume nats message.
type SubscriberItf interface {
	Handle(ctx context.Context, message *nats.Msg) error
}

type PublisherItf interface {
	Publish(subj string, data []byte) error
}
