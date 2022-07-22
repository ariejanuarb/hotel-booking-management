package service

import (
	"booking-hotel/exception"
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"booking-hotel/repository"
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"strconv"
	"time"
)

type BookingServiceImpl struct {
	BookingRepository repository.BookingRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewBookingService(bookingRepository repository.BookingRepository, db *sql.DB, validate *validator.Validate) *BookingServiceImpl {
	return &BookingServiceImpl{
		BookingRepository: bookingRepository,
		DB:                db,
		Validate:          validate,
	}
}

func (service *BookingServiceImpl) Create(ctx context.Context, request *web.BookingCreateRequest) (*web.BookingResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	startTime := request.Event_start
	endTime := request.Event_end
	roomId := request.Room_id

	var price int
	var priceHour string
	var priceDay string
	var days int

	if startTime > time.Now().Format("2006-01-02 15:04:05") {
		if endTime > startTime {
			checkSchadules := service.BookingRepository.FindAll(ctx, tx)
			for _, checkSchadule := range checkSchadules {
				if roomId == checkSchadule.Room_id {
					if checkSchadule.Status == "Booked" {
						if startTime >= checkSchadule.Event_start {
							if startTime <= checkSchadule.Event_end {
								return nil, errors.New("Room is Full")
							}
						} else {
							if endTime >= checkSchadule.Event_start {
								return nil, errors.New("Room is Full")
							}
						}
					}
				}
				checkPrices := service.BookingRepository.CheckPrice(ctx, tx)
				for _, checkPrice := range checkPrices {
					if roomId == checkPrice.Id {
						priceHour = checkPrice.Price_per_hour
						priceDay = checkPrice.Price_per_day
					}
				}
			}
			layout := "2006-01-02T15:04:05Z"

			timeStart, _ := time.Parse(layout, startTime)
			timeEnd, _ := time.Parse(layout, endTime)
			diff := timeEnd.Sub(timeStart).Hours()

			priceHours, _ := strconv.Atoi(priceHour)
			priceDays, _ := strconv.Atoi(priceDay)
			day := time.Hour.Hours() * 24

			if int(diff)%24 == 0 {
				days = int(diff / day)
			}

			if int(diff)%24 == 0 {
				price = priceDays * days
			} else if diff < day {
				price = priceHours * int(diff)
			}
		} else {
			return nil, errors.New("start time must be earlier than end time")
		}
	} else {
		return nil, errors.New("Changes Date!!")

	}

	booking := &domain.Booking{
		Status:             "Booked",
		Room_id:            request.Room_id,
		Hotel_id:           request.Hotel_id,
		Pic_name:           request.Pic_name,
		Pic_Contact:        request.Pic_Contact,
		Event_start:        request.Event_start,
		Event_end:          request.Event_end,
		Invoice_number:     "INV-" + strconv.Itoa(request.Room_id) + "-" + time.Now().Format("20060402150405"),
		Invoice_date:       time.Now(),
		Invoice_grandtotal: strconv.Itoa(price),
		Discount_request:   "Null",
		Created_at:         time.Now(),
		Updated_at:         time.Now(),
	}

	booking = service.BookingRepository.Save(ctx, tx, booking)

	return helper.ToBookingResponse(booking), nil
}

func (service *BookingServiceImpl) Cancel(ctx context.Context, request *web.UpdateRequest) (*web.BookingResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	bookingResponse, err := service.BookingRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	bookings := helper.ToBooking(*bookingResponse)
	layout := "2006-01-02T15:04:05Z"

	startTime := bookings.Event_start
	start, _ := time.Parse(layout, startTime)
	diff := start.Sub(time.Now()).Hours()

	if bookings.Status == "Canceled" {
		return nil, errors.New("This booking is already Canceled")
	}
	if diff >= 164 {
		bookings.Status = "Canceled"
		bookings.Updated_at = time.Now()
		bookings = service.BookingRepository.Cancel(ctx, tx, bookings)
	} else {
		return nil, errors.New("you can't cancel this anymore")
	}

	return helper.ToBookingResponse(bookings), nil
}

func (service *BookingServiceImpl) FindById(ctx context.Context, bookingId int) *web.BookingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	booking, err := service.BookingRepository.FindById(ctx, tx, bookingId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return booking
}

func (service *BookingServiceImpl) FindAll(ctx context.Context) []*web.BookingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	booking := service.BookingRepository.FindAll(ctx, tx)
	return booking
}

func (service *BookingServiceImpl) Discount(ctx context.Context, request *web.Discount) (*web.BookingResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	bookingResponse, err := service.BookingRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	if bookingResponse.Status == "Booked" {
		if bookingResponse.Discount_request != "Null" {
			return nil, errors.New("Request Already Submited")
		}
		bookings := helper.ToBooking(*bookingResponse)
		bookings.Discount_request = "Pending"
		bookings.Updated_at = time.Now()

		bookings = service.BookingRepository.Discount(ctx, tx, bookings)
		return helper.ToBookingResponse(bookings), nil
	} else {
		return nil, errors.New("Booking is already canceled")
	}
}

func (service *BookingServiceImpl) ResponseDiscount(ctx context.Context, request *web.ResponseDiscount) (*web.BookingResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	bookingResponse, err := service.BookingRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	if bookingResponse.Status == "Booked" {
		if bookingResponse.Discount_request == "Pending" {
			bookings := helper.ToBooking(*bookingResponse)
			bookings.Discount_request = request.Discount_request
			bookings.Updated_at = time.Now()

			bookings = service.BookingRepository.ResponseDiscount(ctx, tx, bookings)

			return helper.ToBookingResponse(bookings), nil
		} else {
			return nil, errors.New("There is no request / request already answered")
		}
	} else {
		return nil, errors.New("Booking is already canceled")
	}
}

func (service *BookingServiceImpl) FindAllDiscount(ctx context.Context) []*web.BookingResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	discount := service.BookingRepository.FindAllDiscount(ctx, tx)
	return discount
}

func (service *BookingServiceImpl) HotelRevenue(ctx context.Context, hotelId int, invoiceDate string) *web.HotelRevenueResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	revenue, err := service.BookingRepository.HotelRevenue(ctx, tx, hotelId, invoiceDate)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return revenue
}

func (service *BookingServiceImpl) Revenue(ctx context.Context, invoiceDate string) *web.RevenueResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	revenue, err := service.BookingRepository.Revenue(ctx, tx, invoiceDate)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return revenue
}
