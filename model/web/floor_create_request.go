package web

type FloorCreateRequest struct {
	Number  int `validate:"required" json:"number"`
	HotelId int `validate:"required" json:"hotel_id"`
	RoomId  int `validate:"required" json:"room_id"`
}
