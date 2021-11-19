package job

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/rizalgowandy/gdk/pkg/logx"
)

func NewB(ctx context.Context) message.NoPublishHandlerFunc {
	return func(msg *message.Message) error {
		var payload PayloadMessage

		err := json.Unmarshal(msg.Payload, &payload)
		if err != nil {
			return err
		}

		logx.INF(ctx, payload, "success")

		return nil
	}
}
