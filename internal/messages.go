package internal

import (
	"fmt"
	"strconv"
	"strings"
	"golang.org/x/text/language"
	libMessage "golang.org/x/text/message"
)

type Request struct {
	Asset     string `json:"asset"`
	Fiat      string `json:"fiat"`
	TradeType string `json:"tradeType"`
	Page      int    `json:"page"`
	Rows      int    `json:"rows"`
}

type Response struct {
	Code          string              `json:"code"`
	Message       *string             `json:"message"`
	MessageDetail *string             `json:"messageDetail"`
	Data          []AdvertisementItem `json:"data"`
	Total         int                 `json:"total"`
	Success       bool                `json:"success"`
}

type AdvertisementItem struct {
	Adv           Advertisement `json:"adv"`
	Advertiser    Advertiser    `json:"advertiser"`
	PrivilegeDesc *string       `json:"privilegeDesc"`
	PrivilegeType *string       `json:"privilegeType"`
}

type Advertisement struct {
	AdvNo                         string        `json:"advNo"`
	Classify                      string        `json:"classify"`
	TradeType                     string        `json:"tradeType"`
	Asset                         string        `json:"asset"`
	FiatUnit                      string        `json:"fiatUnit"`
	AdvStatus                     *string       `json:"advStatus"`
	Price                         string        `json:"price"`
	SurplusAmount                 string        `json:"surplusAmount"`
	TradableQuantity              string        `json:"tradableQuantity"`
	MaxSingleTransAmount          string        `json:"maxSingleTransAmount"`
	MinSingleTransAmount          string        `json:"minSingleTransAmount"`
	PayTimeLimit                  int           `json:"payTimeLimit"`
	TradeMethods                  []TradeMethod `json:"tradeMethods"`
	DynamicMaxSingleTransAmount   string        `json:"dynamicMaxSingleTransAmount"`
	MinSingleTransQuantity        string        `json:"minSingleTransQuantity"`
	MaxSingleTransQuantity        string        `json:"maxSingleTransQuantity"`
	DynamicMaxSingleTransQuantity string        `json:"dynamicMaxSingleTransQuantity"`
	CommissionRate                string        `json:"commissionRate"`
	AssetScale                    int           `json:"assetScale"`
	FiatScale                     int           `json:"fiatScale"`
	PriceScale                    int           `json:"priceScale"`
	FiatSymbol                    string        `json:"fiatSymbol"`
	IsTradable                    bool          `json:"isTradable"`
	IsSafePayment                 bool          `json:"isSafePayment"`

	// optional fields
	Remarks       *string `json:"remarks"`
	AutoReplyMsg  *string `json:"autoReplyMsg"`
	InventoryType *string `json:"inventoryType"`
	OfflineReason *string `json:"offlineReason"`
	CloseReason   *string `json:"closeReason"`
}

type TradeMethod struct {
	PayId               *string `json:"payId"`
	PayMethodId         string  `json:"payMethodId"`
	PayType             string  `json:"payType"`
	PayAccount          *string `json:"payAccount"`
	PayBank             *string `json:"payBank"`
	PaySubBank          *string `json:"paySubBank"`
	Identifier          string  `json:"identifier"`
	IconUrlColor        *string `json:"iconUrlColor"`
	TradeMethodName     string  `json:"tradeMethodName"`
	TradeMethodShortName *string `json:"tradeMethodShortName"`
	TradeMethodBgColor  string  `json:"tradeMethodBgColor"`
}

type Advertiser struct {
	UserNo             string    `json:"userNo"`
	RealName           *string   `json:"realName"`
	NickName           string    `json:"nickName"`
	OrderCount         *int      `json:"orderCount"`
	MonthOrderCount    int       `json:"monthOrderCount"`
	MonthFinishRate    float64   `json:"monthFinishRate"`
	PositiveRate       float64   `json:"positiveRate"`
	UserType           string    `json:"userType"`
	UserGrade          int       `json:"userGrade"`
	UserIdentity       string    `json:"userIdentity"`
	ProMerchant        *bool     `json:"proMerchant"`
	Badges             *[]string `json:"badges"`
	VipLevel           *int      `json:"vipLevel"`
	IsBlocked          bool      `json:"isBlocked"`
	ActiveTimeInSecond int       `json:"activeTimeInSecond"`
	TagIconUrls        []string  `json:"tagIconUrls"`
}

// FormatAlertMessage creates a formatted message for WhatsApp alerts
func FormatAlertMessage(ads []AdvertisementItem, config *Config) string {
	if len(ads) == 0 {
		return ""
	}

	printer := libMessage.NewPrinter(language.English)
	var message strings.Builder
	message.WriteString("--------------------------------\n🚨 *Binance P2P Alert*\n--------------------------------\n\n")
	message.WriteString(fmt.Sprintf("Found %d ads matching your criteria:\n\n", len(ads)))
	for index, ad := range ads {
		price, err := strconv.ParseFloat(ad.Adv.Price, 64)
		if err != nil {
			price = 0
		}
		message.WriteString(fmt.Sprintf("💰 Price: %s %s\n", printer.Sprintf("%.2f", price), ad.Adv.FiatUnit))
		minSingleTransAmount, err := strconv.ParseFloat(ad.Adv.MinSingleTransAmount, 64)
		if err != nil {
			minSingleTransAmount = 0
		}
		maxSingleTransAmount, err := strconv.ParseFloat(ad.Adv.MaxSingleTransAmount, 64)
		if err != nil {
			maxSingleTransAmount = 0
		}

		message.WriteString(fmt.Sprintf("💵 Range Amount: %s - %s %s\n", printer.Sprintf("%.2f", minSingleTransAmount), printer.Sprintf("%.2f", maxSingleTransAmount), ad.Adv.FiatUnit))
		message.WriteString(fmt.Sprintf("👤 Trader: %s\n", ad.Advertiser.NickName))
		message.WriteString(fmt.Sprintf("⭐ Positive Rate: %s%%\n\n", printer.Sprintf("%.2f", (ad.Advertiser.PositiveRate * 100))))
		if len(ads) > 3 && index >= 3 {
			break
		}
	}
	message.WriteString(fmt.Sprintf("Check more ads at: https://p2p.binance.com/trade/all-payments/%s?fiat=%s", config.Asset, config.Fiat))
	return message.String()
}