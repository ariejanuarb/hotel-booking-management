package unitTest

import "booking-hotel/model/domain"

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{r: r}
}

func (s *Service) Get(id int) (*domain.Booking, error) {
	return s.r.Get(id)
}

func (s *Service) Create(status string, hotel_id int, room_id int, pic_name string, pic_contact string, event_start string, event_end string, invoice_number string, invoice_grandtotal string, discount_request string) (int, error) {
	booking, err := domain.NewBooking(status, hotel_id, room_id, pic_name, pic_contact, event_start, event_end, invoice_number, invoice_grandtotal, discount_request)
	if err != nil {
		return 0, err
	}
	return s.r.Create(booking)
}

func (s *Service) Update(id int, status string) error {
	booking, err := s.r.Get(id)
	if err != nil {
		return err
	}

	err = booking.Update(status)
	if err != nil {
		return err
	}

	return s.r.Update(booking)
}

func (s *Service) Discount(id int, discount string) error {
	booking, err := s.r.Get(id)
	if err != nil {
		return err
	}

	err = booking.Discount(discount)
	if err != nil {
		return err
	}

	return s.r.Discount(booking)
}

func (s *Service) DiscountResponse(id int, discount string) error {
	booking, err := s.r.Get(id)
	if err != nil {
		return err
	}

	err = booking.DiscountResponse(discount)
	if err != nil {
		return err
	}

	return s.r.DiscountResponse(booking)
}
