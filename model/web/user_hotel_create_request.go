package web

type UserHotelCreateRequest struct {
	UserProfileId int `json:"employee_id"`
	HotelId       int `json:"hotel_id"`
}
