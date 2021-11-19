package job

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/rizalgowandy/gdk/pkg/logx"
	"github.com/rizalgowandy/gdk/pkg/queue/watermillx"
	"github.com/rizalgowandy/gdk/pkg/queue/watermillx/examples/with-amqp/topic"
	"github.com/segmentio/ksuid"
)

func NewA(
	ctx context.Context,
	pub watermillx.Publisher,
) message.NoPublishHandlerFunc {
	return func(msg *message.Message) error {
		ids := []string{
			ksuid.New().String(),
		}

		for _, v := range ids {
			task := BMessage{UserID: v}

			msg, err := watermillx.NewMessage(ctx, task)
			if err != nil {
				continue
			}

			err = pub.Publish(topic.B, msg)
			if err != nil {
				continue
			}

			logx.INF(ctx, task, "success")
		}

		return nil
	}
}
