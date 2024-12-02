package slack_helpers

import (
	"context"
	errors_helpers "demo/helpers/errors"
	"encoding/json"
	"os"
)

type slackPayload struct {
	Text string `json:"text"`
}

type SlackClient struct {
	webhook string
	payload []byte
}

func SetupSlack(text string, ctx context.Context) *SlackClient {
	webhook := os.Getenv("SLACK_WEBHOOK")
	payload := slackPayload{
		Text: text,
	}

	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		errors_helpers.ReportError(ctx, err)
		return nil
	}

	return &SlackClient{
		webhook: webhook,
		payload: payloadBytes,
	}
}
