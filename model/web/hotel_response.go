package web

type HotelResponse struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	Province      string `json:"province"`
	City          string `json:"city"`
	ZipCode       string `json:"zip_code"`
	Star          int    `json:"star"`
	TotalEmployee int    `json:"total_employee"`
	TotalFloor    int    `json:"total_floor"`
	TotalRoom     int    `json:"total_room"`
}
