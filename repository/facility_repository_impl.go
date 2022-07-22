package repository

import (
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"context"
	"database/sql"
	"errors"
)

type FacilityRepositoryImpl struct {
}

func NewFacilityRepository() FacilityRepository {
	return &FacilityRepositoryImpl{}
}

func (repository *FacilityRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, facility domain.Facility) domain.Facility {
	SQL := "insert into facility(type, description) values (?,?)"
	result, err := tx.ExecContext(ctx, SQL, facility.Type, facility.Description)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	facility.Id = int(id)
	return facility
}

func (repository *FacilityRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, facility domain.Facility) domain.Facility {
	SQL := "update facility set Type = ?, Description = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, facility.Type, facility.Description, facility.Id)
	helper.PanicIfError(err)

	return facility
}

func (repository *FacilityRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, facility domain.Facility) {
	SQL := "delete from facility where id = ?"
	_, err := tx.ExecContext(ctx, SQL, facility.Id)
	helper.PanicIfError(err)
}

func (repository *FacilityRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, facilityId int) (domain.Facility, error) {
	SQL := "select id,type,description from facility where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, facilityId)
	helper.PanicIfError(err)
	defer rows.Close()

	facility := domain.Facility{}
	if rows.Next() {
		err := rows.Scan(&facility.Id, &facility.Type, &facility.Description)
		helper.PanicIfError(err)
		return facility, nil
	} else {
		return facility, errors.New("facility is not found")
	}
}

func (repository *FacilityRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) []domain.Facility {
	SQL := "select id, type, description from facility"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var facilitas []domain.Facility
	for rows.Next() {
		facility := domain.Facility{}
		err := rows.Scan(&facility.Id, &facility.Type, &facility.Description)
		helper.PanicIfError(err)
		facilitas = append(facilitas, facility)
	}
	return facilitas
}
