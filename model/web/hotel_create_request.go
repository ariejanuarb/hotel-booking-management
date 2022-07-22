package web

type HotelCreateRequest struct {
	Name     string `validate:"required,min=1,max=100" json:"name"`
	Address  string `validate:"required,min=1,max=100" json:"address"`
	Province string `validate:"required,min=1,max=100" json:"province"`
	City     string `validate:"required,min=1,max=100" json:"city"`
	ZipCode  string `validate:"required,min=1,max=100" json:"zip_code"`
	Star     int    `validate:"required" json:"star"`
}
