package repository

import (
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"context"
	"database/sql"
	"errors"
)

type FloorRepositoryImpl struct {
}

func NewFloorRepository() FloorRepository {
	return &FloorRepositoryImpl{}
}

func (repository *FloorRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, floor domain.Floor) domain.Floor {
	SQL := "insert into floor(number, hotel_id  values (?,?)"
	result, err := tx.ExecContext(ctx, SQL, floor.Number, floor.HotelId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	floor.Id = int(id)
	return floor
}

func (repository *FloorRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, floor domain.Floor) domain.Floor {
	SQL := "update floor set number = ?, hotel_id = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, floor.Number, floor.HotelId, floor.Id)
	helper.PanicIfError(err)

	return floor
}

func (repository *FloorRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, floor domain.Floor) {
	SQL := "delete from floor where id = ?"
	_, err := tx.ExecContext(ctx, SQL, floor.Id)
	helper.PanicIfError(err)
}

func (repository *FloorRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, floorId int) (domain.Floor, error) {
	SQL := "select id, number, hotel_id from floor where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, floorId)
	helper.PanicIfError(err)
	defer rows.Close()

	floor := domain.Floor{}
	if rows.Next() {
		err := rows.Scan(&floor.Id, &floor.Number, &floor.HotelId)
		helper.PanicIfError(err)
		return floor, nil
	} else {
		return floor, errors.New("floor is not found")
	}
}

func (repository *FloorRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Floor {
	SQL := "select id, number, hotel_id from floor"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var floors []domain.Floor
	for rows.Next() {
		floor := domain.Floor{}
		err := rows.Scan(&floor.Id, &floor.Number, &floor.HotelId)
		helper.PanicIfError(err)
		floors = append(floors, floor)
	}
	return floors
}
