package internal

// No imports needed for this file

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