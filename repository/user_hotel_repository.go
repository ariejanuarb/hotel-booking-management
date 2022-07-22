package repository

import (
	"booking-hotel/model/domain"
	"context"
	"database/sql"
)

type UserHotelRepository interface {
	Save(ctx context.Context, tx *sql.Tx, userHotel domain.UserHotel) domain.UserHotel
	Update(ctx context.Context, tx *sql.Tx, userHotel domain.UserHotel) domain.UserHotel
	Delete(ctx context.Context, tx *sql.Tx, userHotel domain.UserHotel)
	FindById(ctx context.Context, tx *sql.Tx, userHotelId int) (domain.UserHotel, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.UserHotel
}
