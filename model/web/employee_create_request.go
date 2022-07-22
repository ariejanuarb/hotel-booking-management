package web

type EmployeeCreateRequest struct {
	Name     string `validate:"required,min=1,max=100" json:"name"`
	Gender   string `validate:"required,min=1,max=100" json:"gender"`
	Email    string `validate:"required,min=1,max=100" json:"email"`
	Password string `validate:"required,min=1,max=100" json:"password"`
}
