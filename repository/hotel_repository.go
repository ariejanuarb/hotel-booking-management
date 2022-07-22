package repository

import (
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"context"
	"database/sql"
)

type HotelRepository interface {
	Save(ctx context.Context, tx *sql.Tx, hotel domain.Hotel) domain.Hotel
	Update(ctx context.Context, tx *sql.Tx, hotel domain.Hotel) domain.Hotel
	Delete(ctx context.Context, tx *sql.Tx, hotel domain.Hotel)
	FindById(ctx context.Context, tx *sql.Tx, hotelId int) (web.HotelResponse, error)
	FindAll(ctx context.Context, tx *sql.Tx) []web.HotelResponse
}
