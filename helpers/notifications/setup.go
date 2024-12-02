package notifications_helpers

import (
	"context"
	slack_helpers "demo/helpers/notifications/slack"
	"errors"
	"os"
)

type Notification interface {
	SendMessage() error
}

func NewNotificationChannel(text string, ctx context.Context) (Notification, error) {
	notification := os.Getenv("NOTIFICATION_CHANNEL")
	switch notification {
	case "slack":
		return slack_helpers.SetupSlack(text, ctx), nil
	default:
		return nil, errors.New("unsupported notification channel")
	}
}
