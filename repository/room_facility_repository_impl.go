package repository

import (
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"context"
	"database/sql"
	"errors"
)

type RoomFacilityRepositoryImpl struct {
}

func NewRoomFacilityRepository() RoomFacilityRepository {
	return &RoomFacilityRepositoryImpl{}
}

func (repository *RoomFacilityRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, roomfacility domain.RoomFacility) domain.RoomFacility {
	SQL := "insert into roomfacility(RoomId,FacilityId) values (?,?)"
	result, err := tx.ExecContext(ctx, SQL, roomfacility.RoomId, roomfacility.FacilityId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	roomfacility.Id = int(id)
	return roomfacility
}

func (repository *RoomFacilityRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, roomfacility domain.RoomFacility) domain.RoomFacility {
	SQL := "update roomfacility set RoomId = ?, FacilityId = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, roomfacility.RoomId, roomfacility.FacilityId, roomfacility.Id)
	helper.PanicIfError(err)

	return roomfacility
}

func (repository *RoomFacilityRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, roomfacility domain.RoomFacility) {
	SQL := "delete from roomfacility where id = ?"
	_, err := tx.ExecContext(ctx, SQL, roomfacility.Id)
	helper.PanicIfError(err)
}

func (repository *RoomFacilityRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, roomfacilityId int) (domain.RoomFacility, error) {
	SQL := "select id,RoomId,FacilityId from roomfacility where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, roomfacilityId)
	helper.PanicIfError(err)
	defer rows.Close()

	roomfacility := domain.RoomFacility{}
	if rows.Next() {
		err := rows.Scan(&roomfacility.Id, &roomfacility.RoomId, &roomfacility.FacilityId)
		helper.PanicIfError(err)
		return roomfacility, nil
	} else {
		return roomfacility, errors.New("roomfacility is not found")
	}
}

func (repository *RoomFacilityRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) []domain.RoomFacility {
	SQL := "select Id,RoomId,FacilityId from roomfacility"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var roomfacilitas []domain.RoomFacility
	for rows.Next() {
		roomfacility := domain.RoomFacility{}
		err := rows.Scan(&roomfacility.Id, &roomfacility.RoomId, &roomfacility.FacilityId)
		helper.PanicIfError(err)
		roomfacilitas = append(roomfacilitas, roomfacility)
	}
	return roomfacilitas
}
