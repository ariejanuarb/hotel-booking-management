package domain

import (
	"errors"
	"math/rand"
	"time"
)

type Booking struct {
	Id                 int
	Status             string
	Room_id            int
	Hotel_id           int
	Pic_name           string
	Pic_Contact        string
	Event_start        string
	Event_end          string
	Invoice_number     string
	Invoice_date       time.Time
	Invoice_grandtotal string
	Discount_request   string
	Created_at         time.Time
	Updated_at         time.Time
}

func NewBooking(status string, room_id int, hotel_id int, pic_name string, pic_contact string, event_start string, event_end string, invoice_number string, invoice_grandtotal string, discount_request string) (*Booking, error) {
	booking := &Booking{
		Id:                 rand.Int(),
		Status:             status,
		Room_id:            room_id,
		Hotel_id:           hotel_id,
		Pic_name:           pic_name,
		Pic_Contact:        pic_contact,
		Event_start:        event_start,
		Event_end:          event_end,
		Invoice_number:     invoice_number,
		Invoice_grandtotal: invoice_grandtotal,
		Discount_request:   discount_request,
	}
	err := booking.Validate()
	if err != nil {
		return booking, err
	}
	return booking, nil
}

func (b *Booking) Validate() error {
	if b.Status == "" || b.Room_id == 0 || b.Hotel_id == 0 || b.Pic_name == "" || b.Pic_Contact == "" || b.Event_start == "" || b.Event_end == "" || b.Invoice_number == "" || b.Invoice_grandtotal == "" || b.Discount_request == "" {
		return errors.New("invalid entity")
	}

	if b.Event_start > b.Event_end {
		return errors.New("event start must be earlier")
	}

	return nil
}

func (b *Booking) Update(status string) error {
	b.Status = status

	return b.Validate()
}

func (b *Booking) Discount(discount string) error {
	b.Discount_request = discount

	return b.Validate()
}

func (b *Booking) DiscountResponse(discountResponse string) error {
	b.Discount_request = discountResponse

	return b.Validate()
}
