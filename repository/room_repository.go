package repository

import (
	"booking-hotel/model/domain"
	"context"
	"database/sql"
)

type RoomRepository interface {
	Save(ctx context.Context, tx *sql.Tx, room domain.Room) domain.Room
	Update(ctx context.Context, tx *sql.Tx, room domain.Room) domain.Room
	Delete(ctx context.Context, tx *sql.Tx, room domain.Room)
	FindById(ctx context.Context, tx *sql.Tx, roomId int) (domain.Room, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []domain.Room
}
