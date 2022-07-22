package web

type RoomResponse struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Capacity     int     `json:"capacity"`
	PricePerHour float64 `json:"price_per_hour"`
	PricePerDay  float64 `json:"price_per_day"`
	Facility_id  int     `json:"facility_id"`
}
