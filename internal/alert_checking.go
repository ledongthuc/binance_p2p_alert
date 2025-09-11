package internal

import (
	"strconv"
)

func CheckConditions(ads *Response, config *Config) []AdvertisementItem {
	adsWithAlert := []AdvertisementItem{}
	for _, ad := range ads.Data {
		price, err := strconv.ParseFloat(ad.Adv.Price, 64)
		if err != nil || price > config.MaxPrice {
			continue
		}

		maxAmount, err := strconv.ParseFloat(ad.Adv.MaxSingleTransAmount, 64)
		if err != nil || maxAmount < config.MinOfMaxAmount {
			continue
		}

		adsWithAlert = append(adsWithAlert, ad)
	}
	return adsWithAlert
}