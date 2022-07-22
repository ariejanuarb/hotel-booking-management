package web

import "time"

type BookingResponse struct {
	Id                 int       `json:"id"`
	Status             string    `json:"status"`
	Room_id            int       `json:"room_Id"`
	Hotel_id           int       `json:"hotel_Id"`
	Price_per_hour     string    `json:"price_Per_Hour"`
	Price_per_day      string    `json:"price_Per_Day"`
	Pic_name           string    `json:"pic_Name"`
	Pic_Contact        string    `json:"pic_Contact"`
	Event_start        string    `json:"event_Start"`
	Event_end          string    `json:"event_End"`
	Invoice_number     string    `json:"invoice_Number"`
	Invoice_date       time.Time `json:"invoice_Date"`
	Invoice_grandtotal string    `json:"invoice_Grandtotal"`
	Discount_request   string    `json:"discount_Request"`
	Created_at         time.Time `json:"created_At"`
	Updated_at         time.Time `json:"updated_At"`
}
