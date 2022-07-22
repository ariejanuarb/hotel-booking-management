package web

type RoomFacilityCreateRequest struct {
	RoomId     string `validate:"required,min=1,max=100" json:"roomId"`
	FacilityId string `validate:"required,min=1,max=100" json:"facilityId"`
}
