package middleware

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

// NewMetadata add identifier information to message metadata.
func NewMetadata(h message.HandlerFunc) message.HandlerFunc {
	return func(msg *message.Message) (events []*message.Message, err error) {
		ctx := msg.Context()
		msg.Metadata.Set("handler_name", message.HandlerNameFromCtx(ctx))
		msg.Metadata.Set("publisher_name", message.PublisherNameFromCtx(ctx))
		msg.Metadata.Set("subscriber_name", message.SubscriberNameFromCtx(ctx))
		msg.Metadata.Set("subscribe_topic", message.SubscribeTopicFromCtx(ctx))
		msg.Metadata.Set("publish_topic", message.PublishTopicFromCtx(ctx))
		return h(msg)
	}
}
