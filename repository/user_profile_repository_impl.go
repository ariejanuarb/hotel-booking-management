package repository

import (
	"booking-hotel/helper"
	"booking-hotel/model/domain"
	"context"
	"database/sql"
	"errors"
)

type UserProfileRepositoryImpl struct {
}

func NewUserProfileRepository() UserProfileRepository {
	return &UserProfileRepositoryImpl{}
}

func (repository *UserProfileRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, userProfile domain.UserProfile) domain.UserProfile {
	SQL := "insert into user_profile(name, gender, email, password, role_id) values (?,?,?,?,1)"
	result, err := tx.ExecContext(ctx, SQL, userProfile.Name, userProfile.Gender, userProfile.Email, userProfile.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	userProfile.Id = int(id)
	return userProfile
}

func (repository *UserProfileRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, userProfile domain.UserProfile) domain.UserProfile {
	SQL := "update user_profile set name = ?, gender = ?, email = ?, password = ?, role_id = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, userProfile.Name, userProfile.Gender, userProfile.Email, userProfile.Password, userProfile.RoleId, userProfile.Id)
	helper.PanicIfError(err)

	return userProfile
}

func (repository *UserProfileRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userProfile domain.UserProfile) {
	SQL := "delete from user_profile where id = ?"
	_, err := tx.ExecContext(ctx, SQL, userProfile.Id)
	helper.PanicIfError(err)
}

func (repository *UserProfileRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userProfileId int) (domain.UserProfile, error) {
	SQL := "select id, name, gender, email, password, role_id from user_profile where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userProfileId)
	helper.PanicIfError(err)
	defer rows.Close()

	userProfile := domain.UserProfile{}
	if rows.Next() {
		err := rows.Scan(&userProfile.Id, &userProfile.Name, &userProfile.Gender, &userProfile.Email, &userProfile.Password, &userProfile.RoleId)
		helper.PanicIfError(err)
		return userProfile, nil
	} else {
		return userProfile, errors.New("userProfile is not found")
	}
}

func (repository *UserProfileRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.UserProfile {
	SQL := "select id, name, gender, email, password, role_id from user_profile"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var userProfiles []domain.UserProfile
	for rows.Next() {
		userProfile := domain.UserProfile{}
		err := rows.Scan(&userProfile.Id, &userProfile.Name, &userProfile.Gender, &userProfile.Email, &userProfile.Password, &userProfile.RoleId)
		helper.PanicIfError(err)
		userProfiles = append(userProfiles, userProfile)
	}
	return userProfiles
}
