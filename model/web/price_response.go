package web

type PriceResponse struct {
	Id             int    `json:"id"`
	Price_per_hour string `json:"price_Per_Hour"`
	Price_per_day  string `json:"price_Per_Day"`
}
