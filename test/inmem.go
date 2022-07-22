package test

import (
	"booking-hotel/model/domain"
	"errors"
)

type inmem struct {
	m map[int]*domain.Booking
}

func newInMem() *inmem {
	var m = map[int]*domain.Booking{}
	return &inmem{m: m}
}

func (m *inmem) Create(booking *domain.Booking) (int, error) {
	m.m[booking.Id] = booking
	return int(booking.Id), nil
}

func (m *inmem) Update(booking *domain.Booking) error {
	_, err := m.Get(booking.Id)
	if err != nil {
		return err
	}
	m.m[booking.Id] = booking
	return nil
}

func (m *inmem) Get(id int) (*domain.Booking, error) {
	newid := int(id)
	if m.m[newid] == nil {
		return nil, errors.New("domain not found")
	}
	return m.m[newid], nil
}

func (m *inmem) Discount(booking *domain.Booking) error {
	_, err := m.Get(booking.Id)
	if err != nil {
		return err
	}
	m.m[booking.Id] = booking
	return nil
}

func (m *inmem) DiscountResponse(booking *domain.Booking) error {
	_, err := m.Get(booking.Id)
	if err != nil {
		return err
	}
	m.m[booking.Id] = booking
	return nil
}
