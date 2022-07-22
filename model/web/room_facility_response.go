package web

type RoomFacilityResponse struct {
	Id         int    `json:"id"`
	RoomId     string `json:"roomId"`
	FacilityId string `json:"facilityId"`
}
