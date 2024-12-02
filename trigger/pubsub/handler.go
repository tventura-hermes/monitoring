package pubsub_trigger

import (
	"context"
	marketplace_domain "demo/api/marketplace/domain"
	"demo/db"
	errors_helpers "demo/helpers/errors"
	notifications_helpers "demo/helpers/notifications"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func handleMessage(msg *pubsub.Message, ctx context.Context) {
	fmt.Println("reading message")

	database, errDb := db.NewDatabase(ctx)

	if errDb != nil {
		errors_helpers.ReportError(ctx, errDb)
		return
	}

	message := marketplace_domain.MessageMongo{
		Data: string(msg.Data),
	}

	err := database.Insert(message, "marketplace")

	if err != nil {
		errors_helpers.ReportError(ctx, err)
	}

	notification, err := notifications_helpers.NewNotificationChannel(string(msg.Data), ctx)

	if err != nil {
		errors_helpers.ReportError(ctx, err)
		return
	}

	errNot := notification.SendMessage()

	if errNot != nil {
		errors_helpers.ReportError(ctx, errNot)
		return
	}

	fmt.Println("Data Inserted")

}
