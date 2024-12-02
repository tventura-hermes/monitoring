package pubsub_trigger

import (
	"context"
	"fmt"
	"os"
	"sync/atomic"

	"cloud.google.com/go/pubsub"
)

func (p *PubsubClient) GetMessage() {
	subId := os.Getenv("PUBSUB_SUBSCRIPTION")
	sub := p.client.Subscription(subId)
	var received int32
	fmt.Println("Listening for Pub/Sub messages...")

	err := sub.Receive(p.ctx, func(ctx context.Context, m *pubsub.Message) {
		atomic.AddInt32(&received, 1)
		handleMessage(m, ctx)
		m.Ack()
	})

	if err != nil {
		fmt.Printf("Error receiving messages: %v\n", err)
	}
}
