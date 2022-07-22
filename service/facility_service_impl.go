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

type FacilityServiceImpl struct {
	FacilityRepository repository.FacilityRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewFacilityService(facilityRepository repository.FacilityRepository, db *sql.DB, validate *validator.Validate) *FacilityServiceImpl {
	return &FacilityServiceImpl{
		FacilityRepository: facilityRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *FacilityServiceImpl) Create(ctx context.Context, request web.FacilityCreateRequest) web.FacilityResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	facility := domain.Facility{
		Type:        request.Type,
		Description: request.Description,
	}

	facility = service.FacilityRepository.Save(ctx, tx, facility)

	return helper.ToFacilityResponse(facility)
}

func (service *FacilityServiceImpl) Update(ctx context.Context, request web.FacilityUpdateRequest) web.FacilityResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	facility, err := service.FacilityRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	facility.Type = request.Type
	facility.Description = request.Description

	facility = service.FacilityRepository.Update(ctx, tx, facility)

	return helper.ToFacilityResponse(facility)
}

func (service *FacilityServiceImpl) Delete(ctx context.Context, facilityId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	facility, err := service.FacilityRepository.FindById(ctx, tx, facilityId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.FacilityRepository.Delete(ctx, tx, facility)
}

func (service *FacilityServiceImpl) FindById(ctx context.Context, facilityId int) web.FacilityResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	facility, err := service.FacilityRepository.FindById(ctx, tx, facilityId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToFacilityResponse(facility)
}

func (service *FacilityServiceImpl) FindByAll(ctx context.Context) []web.FacilityResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	facilitas := service.FacilityRepository.FindByAll(ctx, tx)

	return helper.ToFacilityResponses(facilitas)
}
