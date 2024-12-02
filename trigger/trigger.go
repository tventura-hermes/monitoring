package trigger

import (
	"context"
	pubsub_trigger "demo/trigger/pubsub"
	"errors"
	"os"
)

type Trigger interface {
	GetMessage()
	CloseClient()
}

func NewTrigger(ctx context.Context) (Trigger, error) {
	triggerType := os.Getenv("TRIGGER_TYPE")
	switch triggerType {
	case "pubsub":
		return pubsub_trigger.SetupPubsub(ctx), nil
	default:
		return nil, errors.New("unsupported trigger backend")
	}
}
