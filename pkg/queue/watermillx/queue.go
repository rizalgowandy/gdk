package watermillx

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

//go:generate mockgen -destination=queue_mock.go -package=queue -source=queue.go

type Publisher interface {
	message.Publisher
}

type Subscriber interface {
	message.Subscriber
}
