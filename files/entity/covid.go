package entity

type CovidResponse struct {
	Country        string `json:"country"`
	Cases          int    `json:"cases"`
	TodayCases     int    `json:"todayCases"`
	Deaths         int    `json:"deaths"`
	TodayDeaths    int    `json:"todayDeaths"`
	Recovered      int    `json:"recovered"`
	TodayRecovered int    `json:"todayRecovered"`
	Active         int    `json:"active"`
	Critical       int    `json:"critical"`
}
