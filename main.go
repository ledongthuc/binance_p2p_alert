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
	fmt.Println("Total ads with alert:", len(adsWithAlert))
}
