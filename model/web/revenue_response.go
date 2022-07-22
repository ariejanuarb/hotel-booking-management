package web

type HotelRevenueResponse struct {
	Invoice_grandtotal string `json:"invoice_Grandtotal"`
	Hotel_id           int    `json:"hotel_Id"`
	Name               string `json:"name"`
	Invoice_date       string `json:"invoice_Date"`
}

type RevenueResponse struct {
	Invoice_grandtotal string `json:"invoice_Grandtotal"`
	Invoice_date       string `json:"invoice_Date"`
}
