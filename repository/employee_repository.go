package repository

import (
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"context"
	"database/sql"
)

type EmployeeRepository interface {
	Save(ctx context.Context, tx *sql.Tx, employee domain.Employee) domain.Employee
	Update(ctx context.Context, tx *sql.Tx, employee domain.Employee) domain.Employee
	Delete(ctx context.Context, tx *sql.Tx, employee domain.Employee)
	FindById(ctx context.Context, tx *sql.Tx, employeeId int) (web.EmployeeResponse, error)
	FindAll(ctx context.Context, tx *sql.Tx) []web.EmployeeResponse
}
