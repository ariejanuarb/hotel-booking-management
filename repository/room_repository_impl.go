package repository

import (
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"context"
	"database/sql"
	"errors"
)

type RoomRepositoryImpl struct {
}

func NewRoomRepository() RoomRepository {
	return &RoomRepositoryImpl{}
}

func (repository *RoomRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, room domain.Room) domain.Room {
	SQL := "insert into room(name, capacity, price_per_hour, price_per_day, facility_id) values (?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, room.Name, room.Capacity, room.PricePerHour, room.PricePerDay, room.Facility_id)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	room.Id = int(id)
	return room
}

func (repository *RoomRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, room domain.Room) domain.Room {
	SQL := "update room set name = ?, capacity =?, price_per_hour =?, price_per_day=?, facility_id= ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, room.Id, room.Name, room.Capacity, room.PricePerHour, room.PricePerDay, room.Facility_id)
	helper.PanicIfError(err)

	return room
}

func (repository *RoomRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, room domain.Room) {
	SQL := "delete from room where id = ?"
	_, err := tx.ExecContext(ctx, SQL, room.Id)
	helper.PanicIfError(err)
}

func (repository *RoomRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, roomId int) (domain.Room, error) {
	SQL := "select id, name, capacity, price_per_hour, price_per_day, facility_id from room where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, roomId)
	helper.PanicIfError(err)
	defer rows.Close()

	room := domain.Room{}
	if rows.Next() {
		err := rows.Scan(&room.Id, &room.Name, &room.Capacity, &room.PricePerHour, &room.PricePerDay, &room.Facility_id)
		helper.PanicIfError(err)
		return room, nil
	} else {
		return room, errors.New("room is not found")
	}
}

func (repository *RoomRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) []domain.Room {
	SQL := "select id, name, capacity, price_per_hour, price_per_day, facility_id from room"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var rooms []domain.Room
	for rows.Next() {
		room := domain.Room{}
		err := rows.Scan(&room.Id, &room.Name, &room.Capacity, &room.PricePerHour, &room.PricePerDay, &room.Facility_id)
		helper.PanicIfError(err)
		rooms = append(rooms, room)
	}
	return rooms
}
