package service

import (
	"booking-hotel/exception"
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"booking-hotel/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type HotelServiceImpl struct {
	HotelRepository repository.HotelRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewHotelService(hotelRepository repository.HotelRepository, DB *sql.DB, validate *validator.Validate) HotelService {
	return &HotelServiceImpl{
		HotelRepository: hotelRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *HotelServiceImpl) Create(ctx context.Context, request web.HotelCreateRequest) web.HotelResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	hotel := domain.Hotel{
		Name:     request.Name,
		Address:  request.Address,
		Province: request.Province,
		City:     request.City,
		ZipCode:  request.ZipCode,
		Star:     request.Star,
	}

	hotel = service.HotelRepository.Save(ctx, tx, hotel)

	return helper.ToHotelResponse(hotel)
}

func (service *HotelServiceImpl) Update(ctx context.Context, request web.HotelUpdateRequest) web.HotelResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	hotelResponse, err := service.HotelRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	hotels := helper.ToHotel(hotelResponse)

	hotels.Name = request.Name
	hotels.Address = request.Address
	hotels.Province = request.Province
	hotels.City = request.City
	hotels.ZipCode = request.ZipCode
	hotels.Star = request.Star

	hotels = service.HotelRepository.Update(ctx, tx, hotels)

	return helper.ToHotelResponse(hotels)
}

func (service *HotelServiceImpl) Delete(ctx context.Context, hotelId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	hotelResponse, err := service.HotelRepository.FindById(ctx, tx, hotelId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	hotels := helper.ToHotel(hotelResponse)
	service.HotelRepository.Delete(ctx, tx, hotels)
}

func (service *HotelServiceImpl) FindById(ctx context.Context, hotelId int) web.HotelResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	hotels, err := service.HotelRepository.FindById(ctx, tx, hotelId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return hotels
}

func (service *HotelServiceImpl) FindAll(ctx context.Context) []web.HotelResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	hotels := service.HotelRepository.FindAll(ctx, tx)

	return hotels
}
