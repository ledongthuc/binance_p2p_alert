package internal

import (
	"fmt"
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

	return nil
}