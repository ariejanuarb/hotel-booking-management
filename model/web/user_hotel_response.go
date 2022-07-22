package web

type UserHotelResponse struct {
	Id            int `json:"assignation_number"`
	UserProfileId int `json:"employee_id"`
	HotelId       int `json:"hotel_id"`
}
