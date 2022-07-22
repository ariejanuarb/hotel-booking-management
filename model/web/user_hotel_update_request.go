package web

type UserHotelUpdateRequest struct {
	Id      int `validate:"required"`
	HotelId int `json:"hotel_id"`
}
