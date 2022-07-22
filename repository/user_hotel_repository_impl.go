package repository

import (
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"context"
	"database/sql"
	"errors"
)

type UserHotelRepositoryImpl struct {
}

func NewUserHotelRepository() UserHotelRepository {
	return &UserHotelRepositoryImpl{}
}

func (repository *UserHotelRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, userHotel domain.UserHotel) domain.UserHotel {
	SQL := "insert into user_hotel(user_profile_id, hotel_id) values (?,?)"
	result, err := tx.ExecContext(ctx, SQL, userHotel.UserProfileId, userHotel.HotelId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	userHotel.Id = int(id)
	return userHotel
}

func (repository *UserHotelRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, userHotel domain.UserHotel) domain.UserHotel {
	SQL := "update user_hotel set hotel_id = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, userHotel.HotelId, userHotel.Id)
	helper.PanicIfError(err)

	return userHotel
}

func (repository *UserHotelRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userHotel domain.UserHotel) {
	SQL := "delete from user_hotel where id = ?"
	_, err := tx.ExecContext(ctx, SQL, userHotel.Id)
	helper.PanicIfError(err)
}

func (repository *UserHotelRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userHotelId int) (domain.UserHotel, error) {
	SQL := "select id, user_profile_id, hotel_id from user_hotel where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userHotelId)
	helper.PanicIfError(err)
	defer rows.Close()

	userHotel := domain.UserHotel{}
	if rows.Next() {
		err := rows.Scan(&userHotel.Id, &userHotel.HotelId, &userHotel.UserProfileId)
		helper.PanicIfError(err)
		return userHotel, nil
	} else {
		return userHotel, errors.New("userHotel is not found")
	}
}

func (repository *UserHotelRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.UserHotel {
	SQL := "select id, user_profile_id, hotel_id from user_hotel"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var userHotels []domain.UserHotel
	for rows.Next() {
		userHotel := domain.UserHotel{}
		err := rows.Scan(&userHotel.Id, &userHotel.UserProfileId, &userHotel.HotelId)
		helper.PanicIfError(err)
		userHotels = append(userHotels, userHotel)
	}
	return userHotels
}
