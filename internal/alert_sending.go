package internal

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

// SendAlert sends WhatsApp alerts for matching advertisements
func SendAlert(adsWithAlert []AdvertisementItem, config *Config) error {
	if len(adsWithAlert) == 0 {
		return nil // No alerts to send
	}

	// Format the alert message
	messageBody := FormatAlertMessage(adsWithAlert, config)
	if messageBody == "" {
		return fmt.Errorf("failed to format alert message")
	}

	fmt.Println(messageBody)

	// Get Slack bot token from environment variable for security
	token := os.Getenv("SLACK_BOT_TOKEN")
	if token == "" {
		return fmt.Errorf("Missing SLACK_BOT_TOKEN env var")
	}

	// Initialize Slack client
	api := slack.New(token)

	channelID := os.Getenv("SLACK_CHANNEL_ID")
	if channelID == "" {
		return fmt.Errorf("Missing SLACK_CHANNEL_ID env var")
	}

	// Send message
	channel, timestamp, err := api.PostMessage(
		channelID,
		slack.MsgOptionText(messageBody, false),
	)

	if err != nil {
		return fmt.Errorf("Failed to send message: %v", err)
	}

	fmt.Printf("Message successfully sent to channel %s at %s\n", channel, timestamp)
	return nil
}