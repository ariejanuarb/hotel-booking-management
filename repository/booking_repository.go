package repository

import (
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"context"
	"database/sql"
)

type BookingRepository interface {
	Save(ctx context.Context, tx *sql.Tx, booking *domain.Booking) *domain.Booking
	Cancel(ctx context.Context, tx *sql.Tx, booking *domain.Booking) *domain.Booking
	FindById(ctx context.Context, tx *sql.Tx, bookingId int) (*web.BookingResponse, error)
	FindAll(ctx context.Context, tx *sql.Tx) []*web.BookingResponse
	Discount(ctx context.Context, tx *sql.Tx, booking *domain.Booking) *domain.Booking
	ResponseDiscount(ctx context.Context, tx *sql.Tx, booking *domain.Booking) *domain.Booking
	FindAllDiscount(ctx context.Context, tx *sql.Tx) []*web.BookingResponse
	HotelRevenue(ctx context.Context, tx *sql.Tx, hotelId int, invoiceDate string) (*web.HotelRevenueResponse, error)
	Revenue(ctx context.Context, tx *sql.Tx, invoiceDate string) (*web.RevenueResponse, error)
	CheckPrice(ctx context.Context, tx *sql.Tx) []*web.PriceResponse
}
