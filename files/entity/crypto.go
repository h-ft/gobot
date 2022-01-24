package entity

type CryptoResponse struct {
	Id                string `json:"id"`
	Symbol            string `json:"symbol"`
	MarketCap         string `json:"market_cap_usd"`
	Name              string `json:"name"`
	Price             string `json:"price_usd"`
	PercentChangeHour string `json:"percent_change_1h"`
	PercentChangeDay  string `json:"percent_change_24h"`
	PercentChangeWeek string `json:"percent_change_7d"`
	VolumeDay         string `json:"volume24"`
}
