package watermillx

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/rizalgowandy/gdk/pkg/errorx/v2"
	"github.com/segmentio/ksuid"
)

func NewMessage(_ context.Context, data any) (*message.Message, error) {
	if data == nil {
		return nil, errorx.E("data cannot be empty")
	}

	var payload message.Payload

	switch v := data.(type) {
	case []byte:
		payload = v

	case string:
		payload = []byte(v)

	default:
		res, err := json.Marshal(v)
		if err != nil {
			return nil, errorx.E(err)
		}
		payload = res
	}

	msg := message.NewMessage(
		ksuid.New().String(),
		payload,
	)

	return msg, nil
}
