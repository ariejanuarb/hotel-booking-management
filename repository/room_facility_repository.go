package repository

import (
	"booking-hotel/model/domain"
	"context"
	"database/sql"
)

type RoomFacilityRepository interface {
	Save(ctx context.Context, tx *sql.Tx, roomfacility domain.RoomFacility) domain.RoomFacility
	Update(ctx context.Context, tx *sql.Tx, roomfacility domain.RoomFacility) domain.RoomFacility
	Delete(ctx context.Context, tx *sql.Tx, roomfacility domain.RoomFacility)
	FindById(ctx context.Context, tx *sql.Tx, roomfacilityId int) (domain.RoomFacility, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []domain.RoomFacility
}
