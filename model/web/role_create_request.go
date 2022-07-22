package web

type RoleCreateRequest struct {
	Type string `validate:"required,min=1,max=100" json:"type"`
}
