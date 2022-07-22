package web

type RoomFacilityUpdateRequest struct {
	Id         int    `validate:"required"`
	RoomId     string `validate:"required,max=200,min=1" json:"roomId"`
	FacilityId string `validate:"required,max=200,min=1" json:"facilityId"`
}
