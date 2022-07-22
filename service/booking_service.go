package service

import (
	"booking-hotel/model/web"
	"context"
)

type BookingService interface {
	Create(ctx context.Context, request *web.BookingCreateRequest) (*web.BookingResponse, error)
	Cancel(ctx context.Context, request *web.UpdateRequest) (*web.BookingResponse, error)
	FindById(ctx context.Context, bookingId int) *web.BookingResponse
	FindAll(ctx context.Context) []*web.BookingResponse
	Discount(ctx context.Context, request *web.Discount) (*web.BookingResponse, error)
	ResponseDiscount(ctx context.Context, request *web.ResponseDiscount) (*web.BookingResponse, error)
	FindAllDiscount(ctx context.Context) []*web.BookingResponse
	HotelRevenue(ctx context.Context, hotelId int, invoiceDate string) *web.HotelRevenueResponse
	Revenue(ctx context.Context, invoiceDate string) *web.RevenueResponse
}
