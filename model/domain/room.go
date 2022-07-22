package domain

type Room struct {
	Id           int
	Name         string
	Capacity     int
	PricePerHour float64
	PricePerDay  float64
	Facility_id  int
}
