package web

type RoomUpdateRequest struct {
	Id           int     `validate:"required"`
	Name         string  `validate:"required" json:"name"`
	Capacity     int     `validate:"required" json:"capacity"`
	PricePerHour float64 `validate:"required" json:"price_per_hour"`
	PricePerDay  float64 `validate:"required" json:"price_per_day"`
	Facility_id  int     `validate:"required" json:"facility_id"`
}
