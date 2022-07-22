package web

type RoleUpdateRequest struct {
	Id   int    `validate:"required"`
	Type string `validate:"required,max=100,min=1" json:"type"`
}
