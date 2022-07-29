package repository

import (
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"booking-hotel/model/web"
	"context"
	"database/sql"
	"errors"
)

type EmployeeRepositoryImpl struct {
}

func NewEmployeeRepository() EmployeeRepository {
	return &EmployeeRepositoryImpl{}
}

func (repository *EmployeeRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, employee domain.Employee) domain.Employee {
	SQL := "insert into user_profile(name, gender, email, password, role_id) values (?,?,?,?,2)"
	result, err := tx.ExecContext(ctx, SQL, employee.Name, employee.Gender, employee.Email, employee.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	employee.Id = int(id)
	return employee
}

func (repository *EmployeeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, employee domain.Employee) domain.Employee {
	SQL := "update user_profile up set up.name = ?, up.gender = ?, up.email = ?, up.password = ? where up.id = ? and up.role_id = 2"
	_, err := tx.ExecContext(ctx, SQL, employee.Name, employee.Gender, employee.Email, employee.Password, employee.Id)
	helper.PanicIfError(err)

	return employee
}

func (repository *EmployeeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, employee domain.Employee) {
	SQL := "delete from user_profile where id = ?"
	_, err := tx.ExecContext(ctx, SQL, employee.Id)
	helper.PanicIfError(err)
}

func (repository *EmployeeRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, employeeId int) (web.EmployeeResponse, error) {
	SQL := "select up.id, up.name, up.gender, up.email, up.password, ifnull(user_hotel.hotel_id,0) from user_profile up left join user_hotel on up.id = user_hotel.user_profile_id where up.id = ? and role_id = 2"
	rows, err := tx.QueryContext(ctx, SQL, employeeId)
	helper.PanicIfError(err)
	defer rows.Close()

	employee := web.EmployeeResponse{}
	if rows.Next() {
		err := rows.Scan(&employee.Id, &employee.Name, &employee.Gender, &employee.Email, &employee.Password, &employee.HotelId)
		helper.PanicIfError(err)
		return employee, nil
	} else {
		return employee, errors.New("employee is not found")
	}
}

func (repository *EmployeeRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []web.EmployeeResponse {
	SQL := "select up.id, up.name, up.gender, up.email, up.password, ifnull(user_hotel.hotel_id,0) from user_profile up left join user_hotel on up.id = user_hotel.user_profile_id where up.role_id = 2"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var employees []web.EmployeeResponse
	for rows.Next() {
		employee := web.EmployeeResponse{}
		err := rows.Scan(&employee.Id, &employee.Name, &employee.Gender, &employee.Email, &employee.Password, &employee.HotelId)
		helper.PanicIfError(err)
		employees = append(employees, employee)
	}
	return employees
}
