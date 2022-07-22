package service

import (
	"booking-hotel/model/web"
	"context"
)

type EmployeeService interface {
	Create(ctx context.Context, request web.EmployeeCreateRequest) web.EmployeeResponse
	Update(ctx context.Context, request web.EmployeeUpdateRequest) web.EmployeeResponse
	Delete(ctx context.Context, employeeId int)
	FindById(ctx context.Context, employeeId int) web.EmployeeResponse
	FindAll(ctx context.Context) []web.EmployeeResponse
}
