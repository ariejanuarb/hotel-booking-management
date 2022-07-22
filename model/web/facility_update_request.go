package web

type FacilityUpdateRequest struct {
	Id          int    `validate:"required"`
	Type        string `validate:"required,max=200,min=1" json:"type"`
	Description string `validate:"required,max=200,min=1" json:"description"`
}
