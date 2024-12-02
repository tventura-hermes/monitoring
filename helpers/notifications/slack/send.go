package slack_helpers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func (s SlackClient) SendMessage() error {
	response, err := http.Post(s.webhook, "application/json", bytes.NewBuffer(s.payload))
	if err != nil {
		return fmt.Errorf("error enviando mensaje a Slack: %w", err)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("slack API error: %s - %s", response.Status, string(body))
	}

	return nil
}
