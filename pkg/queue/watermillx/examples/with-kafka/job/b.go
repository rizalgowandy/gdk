package job

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/rizalgowandy/gdk/pkg/logx"
)

type BMessage struct {
	UserID string `json:"user_id"`
}

func NewB(ctx context.Context) message.NoPublishHandlerFunc {
	return func(msg *message.Message) error {
		var payload BMessage

		err := json.Unmarshal(msg.Payload, &payload)
		if err != nil {
			return err
		}

		logx.INF(ctx, payload, "success")

		return nil
	}
}
