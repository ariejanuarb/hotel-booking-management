package web

type EmployeeUpdateRequest struct {
	Id       int    `validate:"required" json:"employee_id"`
	Name     string `validate:"required,min=1,max=100" json:"name"`
	Gender   string `validate:"required,min=1,max=100" json:"gender"`
	Email    string `validate:"required,min=1,max=100" json:"email"`
	Password string `validate:"required,min=1,max=100" json:"password"`
	//HotelId  int    `validate:"required" json:"assigned_to_hotel"`
}
