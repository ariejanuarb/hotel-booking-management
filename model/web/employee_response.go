package web

type EmployeeResponse struct {
	Id       int    `json:"employee_id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
	Password string `json:"password"`
	HotelId  int    `json:"assigned_to_hotel"`
}
