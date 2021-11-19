package main

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/rizalgowandy/gdk/pkg/queue/watermillx"
	"github.com/rizalgowandy/gdk/pkg/queue/watermillx/examples/with-amqp/topic"
	"github.com/rizalgowandy/gdk/pkg/queue/watermillx/examples/with-kafka/job"
	"github.com/rizalgowandy/gdk/pkg/queue/watermillx/pubsub"
	"github.com/stretchr/testify/assert"
)

var (
	integration bool
	client      watermillx.Publisher
)

func TestMain(m *testing.M) {
	flag.BoolVar(&integration, "integration", false, "enable integration test")
	flag.Parse()

	if !integration {
		os.Exit(m.Run())
	}

	logger, err := watermillx.NewLogger()
	if err != nil {
		os.Exit(1)
	}

	address := []string{"kafka:9092"}

	client, err = pubsub.NewKafkaPublisher(logger, address)
	if err != nil {
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestNewA_Integration(t *testing.T) {
	if !integration {
		return
	}

	msg, err := watermillx.NewMessage(context.Background(), struct{}{})
	assert.NoError(t, err)

	err = client.Publish(topic.A, msg)
	assert.NoError(t, err)
}

func TestNewB_Integration(t *testing.T) {
	if !integration {
		return
	}

	msg, err := watermillx.NewMessage(context.Background(), job.BMessage{
		UserID: "random_id",
	})
	assert.NoError(t, err)

	err = client.Publish(topic.B, msg)
	assert.NoError(t, err)
}
