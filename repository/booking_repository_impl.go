package repository

import (
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"context"
	"database/sql"
	"errors"
	"time"
)

type BookingRepositoryImpl struct {
}

func NewBookingRepository() BookingRepository {
	return &BookingRepositoryImpl{}
}

func (repository *BookingRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, booking *domain.Booking) *domain.Booking {
	SQL := "insert into booking(status, room_id, hotel_id, pic_name, pic_contact, event_start, event_end, invoice_number, invoice_date, invoice_grandtotal, discount_request,  created_at, updated_at) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, booking.Status, booking.Room_id, booking.Hotel_id, booking.Pic_name, booking.Pic_Contact, booking.Event_start, booking.Event_end, booking.Invoice_number, time.Now().Format("2006-01-02"), booking.Invoice_grandtotal, booking.Discount_request, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05"))
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	booking.Id = int(id)
	return booking
}

func (repository *BookingRepositoryImpl) Cancel(ctx context.Context, tx *sql.Tx, booking *domain.Booking) *domain.Booking {
	SQL := "update booking set status = ?, updated_at= ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, booking.Status, time.Now().Format("2006-01-02 15:04:05"), booking.Id)
	helper.PanicIfError(err)

	return booking
}

func (repository *BookingRepositoryImpl) Discount(ctx context.Context, tx *sql.Tx, booking *domain.Booking) *domain.Booking {
	SQL := "update booking set discount_request = ?, updated_at= ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, booking.Discount_request, time.Now().Format("2006-01-02 15:04:05"), booking.Id)
	helper.PanicIfError(err)

	return booking
}

func (repository *BookingRepositoryImpl) ResponseDiscount(ctx context.Context, tx *sql.Tx, booking *domain.Booking) *domain.Booking {
	SQL := "update booking set discount_request = ?, updated_at= ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, booking.Discount_request, time.Now().Format("2006-01-02 15:04:05"), booking.Id)
	helper.PanicIfError(err)

	return booking
}

func (repository *BookingRepositoryImpl) FindAllDiscount(ctx context.Context, tx *sql.Tx) []*web.BookingResponse {
	SQL := "select id, discount_request from booking where discount_request = 'Pending' "
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var bookings []*web.BookingResponse
	for rows.Next() {
		booking := &web.BookingResponse{}
		err := rows.Scan(&booking.Id, &booking.Discount_request)
		helper.PanicIfError(err)
		bookings = append(bookings, booking)
	}
	return bookings
}

func (repository *BookingRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bookingId int) (*web.BookingResponse, error) {
	SQL := "select b.id, b.status, b.room_id, b.hotel_id, b.pic_name, b.pic_contact, b.event_start, b.event_end, b.invoice_number, b.invoice_date, b.invoice_grandtotal, b.discount_request, r.price_per_hour, r.price_per_day, b.created_at, b.updated_at from booking b inner join room r on b.room_id=r.id where b.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, bookingId)
	helper.PanicIfError(err)
	defer rows.Close()

	booking := web.BookingResponse{}
	if rows.Next() {
		err := rows.Scan(&booking.Id, &booking.Status, &booking.Room_id, &booking.Hotel_id, &booking.Pic_name, &booking.Pic_Contact, &booking.Event_start, &booking.Event_end, &booking.Invoice_number, &booking.Invoice_date, &booking.Invoice_grandtotal, &booking.Discount_request, &booking.Price_per_hour, &booking.Price_per_day, &booking.Created_at, &booking.Updated_at)
		helper.PanicIfError(err)
		defer rows.Close()
		return &booking, nil
	} else {
		return &booking, errors.New("Booking is not found")
	}
}

func (repository *BookingRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []*web.BookingResponse {
	SQL := "select b.id, b.status, b.room_id, b.hotel_id, b.pic_name, b.pic_contact, b.event_start, b.event_end, b.invoice_number, b.invoice_date, b.invoice_grandtotal, b.discount_request, r.price_per_hour, r.price_per_day, b.created_at, b.updated_at from booking b inner join room r on b.room_id=r.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var bookings []*web.BookingResponse
	for rows.Next() {
		booking := &web.BookingResponse{}
		err := rows.Scan(&booking.Id, &booking.Status, &booking.Room_id, &booking.Hotel_id, &booking.Pic_name, &booking.Pic_Contact, &booking.Event_start, &booking.Event_end, &booking.Invoice_number, &booking.Invoice_date, &booking.Invoice_grandtotal, &booking.Discount_request, &booking.Price_per_hour, &booking.Price_per_day, &booking.Created_at, &booking.Updated_at)
		helper.PanicIfError(err)
		bookings = append(bookings, booking)
	}
	return bookings
}

func (repository *BookingRepositoryImpl) HotelRevenue(ctx context.Context, tx *sql.Tx, hotelId int, invoiceDate string) (*web.HotelRevenueResponse, error) {
	SQL := "select h.id, SUM(b.invoice_grandtotal) as grand_total, h.name, b.invoice_date from booking b inner join hotel h on b.hotel_id=h.id where status = 'Booked' AND h.id =? AND invoice_date = ?"
	rows, err := tx.QueryContext(ctx, SQL, hotelId, invoiceDate)
	helper.PanicIfError(err)
	defer rows.Close()

	revenue := web.HotelRevenueResponse{}
	if rows.Next() {
		err := rows.Scan(&revenue.Hotel_id, &revenue.Invoice_grandtotal, &revenue.Name, &revenue.Invoice_date)
		helper.PanicIfError(err)
		defer rows.Close()
		return &revenue, nil
	} else {
		return &revenue, errors.New("Revenue is not found")
	}
}

func (repository *BookingRepositoryImpl) Revenue(ctx context.Context, tx *sql.Tx, invoiceDate string) (*web.RevenueResponse, error) {
	SQL := "select SUM(b.invoice_grandtotal) as grand_total, b.invoice_date from booking b inner join hotel h on b.hotel_id=h.id where status = 'Booked' AND b.invoice_date = ?"
	rows, err := tx.QueryContext(ctx, SQL, invoiceDate)
	helper.PanicIfError(err)
	defer rows.Close()

	revenue := web.RevenueResponse{}
	if rows.Next() {
		err := rows.Scan(&revenue.Invoice_grandtotal, &revenue.Invoice_date)
		helper.PanicIfError(err)
		defer rows.Close()
		return &revenue, nil
	} else {
		return &revenue, errors.New("Revenue is not found")
	}
}

func (repository *BookingRepositoryImpl) CheckPrice(ctx context.Context, tx *sql.Tx) []*web.PriceResponse {
	SQL := "SELECT id, price_per_hour, price_per_day FROM room"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var rooms []*web.PriceResponse
	for rows.Next() {
		room := &web.PriceResponse{}
		err := rows.Scan(&room.Id, &room.Price_per_hour, &room.Price_per_day)
		helper.PanicIfError(err)
		rooms = append(rooms, room)
	}
	return rooms
}
