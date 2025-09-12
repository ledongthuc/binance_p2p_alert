package main

import (
	"fmt"

	"github.com/ledongthuc/binance_p2p_alert/internal"
)

func main() {
	// Load configuration
	config, err := internal.LoadConfig("config.yaml")
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	ads, err := internal.GetBinanceP2PAds(config)
	if err != nil {
		fmt.Println("Error getting Binance P2P ads:", err)
		return
	}
	adsWithAlert := internal.CheckConditions(ads, config)

	messageBody := internal.FormatAlertMessage(adsWithAlert, config)
	if messageBody == "" {
		return
	}
	fmt.Println(messageBody)
	
	// Send WhatsApp alert if there are matching ads
	if config.EnableSlackAlert && len(adsWithAlert) > 0 {
		err = internal.SendAlert(messageBody)
		if err != nil {
			fmt.Println("Error sending WhatsApp alert:", err)
			return
		}
	} else {
		fmt.Println("No ads matching criteria found.")
	}
}
