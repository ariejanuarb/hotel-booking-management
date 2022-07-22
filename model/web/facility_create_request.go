package web

type FacilityCreateRequest struct {
	Type        string `validate:"required,min=1,max=100" json:"type"`
	Description string `validate:"required,min=1,max=100" json:"description"`
}
