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

type RoomServiceImpl struct {
	RoomRepository repository.RoomRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewRoomService(roomRepository repository.RoomRepository, DB *sql.DB, validate *validator.Validate) RoomService {
	return &RoomServiceImpl{
		RoomRepository: roomRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *RoomServiceImpl) Create(ctx context.Context, request web.RoomCreateRequest) web.RoomResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	room := domain.Room{
		Name:         request.Name,
		Capacity:     request.Capacity,
		PricePerHour: request.PricePerHour,
		PricePerDay:  request.PricePerDay,
		Facility_id:  request.Facility_id,
	}

	room = service.RoomRepository.Save(ctx, tx, room)

	return helper.ToRoomResponse(room)
}

func (service *RoomServiceImpl) Update(ctx context.Context, request web.RoomUpdateRequest) web.RoomResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	room, err := service.RoomRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	room.Name = request.Name
	room.Capacity = request.Capacity
	room.PricePerHour = request.PricePerHour
	room.PricePerDay = request.PricePerDay
	room.Facility_id = request.Facility_id

	room = service.RoomRepository.Update(ctx, tx, room)

	return helper.ToRoomResponse(room)
}

func (service *RoomServiceImpl) Delete(ctx context.Context, roomId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	room, err := service.RoomRepository.FindById(ctx, tx, roomId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.RoomRepository.Delete(ctx, tx, room)
}

func (service *RoomServiceImpl) FindById(ctx context.Context, roomId int) web.RoomResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	room, err := service.RoomRepository.FindById(ctx, tx, roomId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToRoomResponse(room)
}

func (service *RoomServiceImpl) FindByAll(ctx context.Context) []web.RoomResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	rooms := service.RoomRepository.FindByAll(ctx, tx)

	return helper.ToRoomResponses(rooms)
}
