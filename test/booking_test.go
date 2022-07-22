package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Create(t *testing.T) {
	inmem := newInMem()
	service := NewService(inmem)
	_, err := service.Create("Booked", 1, 1, "Rio", "081290314494", "2022-12-12 10:00:00", "2022-12-12 12:00:00", "inv-2367486", "800000", "Null")
	assert.Nil(t, err)
}

func TestService_Get(t *testing.T) {
	inmem := newInMem()
	service := NewService(inmem)
	id, _ := service.Create("Booked", 1, 1, "Rio", "081290314494", "2022-12-12 10:00:00", "2022-12-12 12:00:00", "inv-2367486", "800000", "Null")

	booking, err := service.Get(id)
	assert.Nil(t, err)
	assert.Equal(t, "Booked", booking.Status)
}

func TestService_Update(t *testing.T) {
	inmem := newInMem()
	service := NewService(inmem)
	id, _ := service.Create("Booked", 1, 1, "Rio", "081290314494", "2022-12-12 10:00:00", "2022-12-12 12:00:00", "inv-2367486", "800000", "Null")

	err := service.Update(id, "Canceled")
	assert.Nil(t, err)
}

func TestService_Discount(t *testing.T) {
	inmem := newInMem()
	service := NewService(inmem)
	id, _ := service.Create("Booked", 1, 1, "Rio", "081290314494", "2022-12-12 10:00:00", "2022-12-12 12:00:00", "inv-2367486", "800000", "Null")

	err := service.Discount(id, "Requested")
	assert.Nil(t, err)
}

func TestService_DiscountResponse(t *testing.T) {
	inmem := newInMem()
	service := NewService(inmem)
	id, _ := service.Create("Booked", 1, 1, "Rio", "081290314494", "2022-12-12 10:00:00", "2022-12-12 12:00:00", "inv-2367486", "800000", "Null")

	err := service.DiscountResponse(id, "Accepted")
	assert.Nil(t, err)
}
