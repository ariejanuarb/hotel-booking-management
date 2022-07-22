package web

type HotelUpdateRequest struct {
	Id       int    `validate:"required"`
	Name     string `validate:"required,max=100,min=1" json:"name"`
	Address  string `validate:"required,max=100,min=1" json:"address"`
	Province string `validate:"required,max=100,min=1" json:"province"`
	City     string `validate:"required,max=100,min=1" json:"city"`
	ZipCode  string `validate:"required,max=100,min=1" json:"zip_code"`
	Star     int    `validate:"required" json:"star"`
}
