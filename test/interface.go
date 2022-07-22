package test

import "booking-hotel/model/domain"

type Writer interface {
	Create(booking *domain.Booking) (int, error)
	Update(booking *domain.Booking) error
	Discount(booking *domain.Booking) error
	DiscountResponse(booking *domain.Booking) error
}

type Reader interface {
	Get(id int) (*domain.Booking, error)
}

type Repository interface {
	Writer
	Reader
}
