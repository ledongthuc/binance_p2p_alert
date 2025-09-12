package internal

import (
	"fmt"
	"strconv"
)

func CheckConditions(ads *Response, config *Config) []AdvertisementItem {
	adsWithAlert := []AdvertisementItem{}
	for _, ad := range ads.Data {
		if config.MaxPrice != nil {
			price, err := strconv.ParseFloat(ad.Adv.Price, 64)
			if err != nil || price > *config.MaxPrice {
				fmt.Println("price", price, "maxPrice", *config.MaxPrice)
				continue
			}
		}

		if config.MinOfMaxAmount != nil {
			maxAmount, err := strconv.ParseFloat(ad.Adv.MaxSingleTransAmount, 64)
			if err != nil || maxAmount < *config.MinOfMaxAmount {
				fmt.Println("maxAmount", maxAmount, "minOfMaxAmount", *config.MinOfMaxAmount)
				continue
			}
		}

		if config.MinRating != nil {
			if ad.Advertiser.PositiveRate < (*config.MinRating / 100) {
				continue
			}
		}

		adsWithAlert = append(adsWithAlert, ad)
	}
	return adsWithAlert
}