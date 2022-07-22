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

type EmployeeServiceImpl struct {
	EmployeeRepository repository.EmployeeRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewEmployeeService(employeeRepository repository.EmployeeRepository, DB *sql.DB, validate *validator.Validate) EmployeeService {
	return &EmployeeServiceImpl{
		EmployeeRepository: employeeRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *EmployeeServiceImpl) Create(ctx context.Context, request web.EmployeeCreateRequest) web.EmployeeResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	employee := domain.Employee{
		Name:     request.Name,
		Gender:   request.Gender,
		Email:    request.Email,
		Password: request.Password,
	}

	employee = service.EmployeeRepository.Save(ctx, tx, employee)

	return helper.ToEmployeeResponse(employee)
}

func (service *EmployeeServiceImpl) Update(ctx context.Context, request web.EmployeeUpdateRequest) web.EmployeeResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	employeeResponse, err := service.EmployeeRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	employe := helper.ToEmployee(employeeResponse)

	employe.Name = request.Name
	employe.Gender = request.Gender
	employe.Email = request.Email
	employe.Password = request.Password

	employe = service.EmployeeRepository.Update(ctx, tx, employe)

	return helper.ToEmployeeResponse(employe)
}

func (service *EmployeeServiceImpl) Delete(ctx context.Context, employeeId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	employeeResponse, err := service.EmployeeRepository.FindById(ctx, tx, employeeId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	employe := helper.ToEmployee(employeeResponse)
	service.EmployeeRepository.Delete(ctx, tx, employe)
}

func (service *EmployeeServiceImpl) FindById(ctx context.Context, employeeId int) web.EmployeeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	employe, err := service.EmployeeRepository.FindById(ctx, tx, employeeId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return employe
}

func (service *EmployeeServiceImpl) FindAll(ctx context.Context) []web.EmployeeResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	employees := service.EmployeeRepository.FindAll(ctx, tx)

	return employees
}
