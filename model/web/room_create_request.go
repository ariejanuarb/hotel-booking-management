package web

type RoomCreateRequest struct {
	Name         string  `validate:"required,min=1,max=100" json:"name"`
	Capacity     int     `validate:"required" json:"capacity"`
	PricePerHour float64 `validate:"required" json:"price_per_hour"`
	PricePerDay  float64 `validate:"required" json:"price_per_day"`
	Facility_id  int     `validate:"required" json:"facility_id"`
}
