package repository

import (
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"context"
	"database/sql"
	"errors"
)

type RoleRepositoryImpl struct {
}

func NewRoleRepository() RoleRepository {
	return &RoleRepositoryImpl{}
}

func (repository *RoleRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role {
	SQL := "insert into role(type) values (?)"
	result, err := tx.ExecContext(ctx, SQL, role.Type)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	role.Id = int(id)
	return role
}

func (repository *RoleRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role {
	SQL := "update role set type = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, role.Type, role.Id)
	helper.PanicIfError(err)

	return role
}

func (repository *RoleRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, role domain.Role) {
	SQL := "delete from role where id = ?"
	_, err := tx.ExecContext(ctx, SQL, role.Id)
	helper.PanicIfError(err)
}

func (repository *RoleRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, roleId int) (domain.Role, error) {
	SQL := "select id, type from role where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, roleId)
	helper.PanicIfError(err)
	defer rows.Close()

	role := domain.Role{}
	if rows.Next() {
		err := rows.Scan(&role.Id, &role.Type)
		helper.PanicIfError(err)
		return role, nil
	} else {
		return role, errors.New("role is not found")
	}
}

func (repository *RoleRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Role {
	SQL := "select id, type from role"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var roles []domain.Role
	for rows.Next() {
		role := domain.Role{}
		err := rows.Scan(&role.Id, &role.Type)
		helper.PanicIfError(err)
		roles = append(roles, role)
	}
	return roles
}
