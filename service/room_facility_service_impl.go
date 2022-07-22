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

type RoomFacilityServiceImpl struct {
	RoomFacilityRepository repository.RoomFacilityRepository
	DB                     *sql.DB
	Validate               *validator.Validate
}

func NewRoomFacilityService(roomfacilityRepository repository.RoomFacilityRepository, db *sql.DB, validate *validator.Validate) *RoomFacilityServiceImpl {
	return &RoomFacilityServiceImpl{
		RoomFacilityRepository: roomfacilityRepository,
		DB:                     db,
		Validate:               validate,
	}
}

func (service *RoomFacilityServiceImpl) Create(ctx context.Context, request web.RoomFacilityCreateRequest) web.RoomFacilityResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	roomfacility := domain.RoomFacility{
		RoomId:     request.RoomId,
		FacilityId: request.FacilityId,
	}

	roomfacility = service.RoomFacilityRepository.Save(ctx, tx, roomfacility)

	return helper.ToRoomFacilityResponse(roomfacility)
}

func (service *RoomFacilityServiceImpl) Update(ctx context.Context, request web.RoomFacilityUpdateRequest) web.RoomFacilityResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	roomfacility, err := service.RoomFacilityRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	roomfacility.RoomId = request.RoomId
	roomfacility.FacilityId = request.FacilityId

	roomfacility = service.RoomFacilityRepository.Update(ctx, tx, roomfacility)

	return helper.ToRoomFacilityResponse(roomfacility)
}

func (service *RoomFacilityServiceImpl) Delete(ctx context.Context, roomfacilityId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	roomfacility, err := service.RoomFacilityRepository.FindById(ctx, tx, roomfacilityId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.RoomFacilityRepository.Delete(ctx, tx, roomfacility)
}

func (service *RoomFacilityServiceImpl) FindById(ctx context.Context, roomfacilityId int) web.RoomFacilityResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	roomfacility, err := service.RoomFacilityRepository.FindById(ctx, tx, roomfacilityId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToRoomFacilityResponse(roomfacility)
}

func (service *RoomFacilityServiceImpl) FindByAll(ctx context.Context) []web.RoomFacilityResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	roomfacilitas := service.RoomFacilityRepository.FindByAll(ctx, tx)

	return helper.ToRoomFacilityResponses(roomfacilitas)
}
