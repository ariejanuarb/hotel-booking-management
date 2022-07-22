package repository

import (
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"context"
	"database/sql"
	"errors"
)

type HotelRepositoryImpl struct {
}

func NewHotelRepository() HotelRepository {
	return &HotelRepositoryImpl{}
}

func (repository *HotelRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, hotel domain.Hotel) domain.Hotel {
	SQL := "insert into hotel(name, address, province, city, zip_code, star) values (?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, hotel.Name, hotel.Address, hotel.Province, hotel.City, hotel.ZipCode, hotel.Star)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	hotel.Id = int(id)
	return hotel
}

func (repository *HotelRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, hotel domain.Hotel) domain.Hotel {
	SQL := "update hotel set name = ?, address = ?, province = ?, city = ?, zip_code = ?, star = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, hotel.Name, hotel.Address, hotel.Province, hotel.City, hotel.ZipCode, hotel.Star, hotel.Id)
	helper.PanicIfError(err)

	return hotel
}

func (repository *HotelRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, hotel domain.Hotel) {
	SQL := "delete from hotel where id = ?"
	_, err := tx.ExecContext(ctx, SQL, hotel.Id)
	helper.PanicIfError(err)
}

func (repository *HotelRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, hotelId int) (web.HotelResponse, error) {
	SQL := "select h.id, h.name, h.address, h.province, h.city, h.zip_code, h.star, (select COUNT(uh.id) from user_hotel uh where uh.hotel_id = h.id) as total_employee, (select COUNT(DISTINCT(f.number)) from floor f where f.hotel_id = h.id), (select COUNT(f.room_id) from floor f where f.hotel_id = h.id) from hotel h where h.id = ? group by h.id"
	rows, err := tx.QueryContext(ctx, SQL, hotelId)
	helper.PanicIfError(err)
	defer rows.Close()

	hotel := web.HotelResponse{}
	if rows.Next() {
		err := rows.Scan(&hotel.Id, &hotel.Name, &hotel.Address, &hotel.Province, &hotel.City, &hotel.ZipCode, &hotel.Star, &hotel.TotalEmployee, &hotel.TotalFloor, &hotel.TotalRoom)
		helper.PanicIfError(err)
		return hotel, nil
	} else {
		return hotel, errors.New("hotel is not found")
	}
}

func (repository *HotelRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []web.HotelResponse {
	SQL := "select h.id, h.name, h.address, h.province, h.city, h.zip_code, h.star, (select COUNT(uh.id) from user_hotel uh where uh.hotel_id = h.id) as total_employee, (select COUNT(DISTINCT(f.number)) from floor f where f.hotel_id = h.id) as total_floor, (select COUNT(f.room_id) from floor f where f.hotel_id = h.id) as total_room from hotel h group by h.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var hotels []web.HotelResponse
	for rows.Next() {
		hotel := web.HotelResponse{}
		err := rows.Scan(&hotel.Id, &hotel.Name, &hotel.Address, &hotel.Province, &hotel.City, &hotel.ZipCode, &hotel.Star, &hotel.TotalEmployee, &hotel.TotalFloor, &hotel.TotalRoom)
		helper.PanicIfError(err)
		hotels = append(hotels, hotel)
	}
	return hotels
}
