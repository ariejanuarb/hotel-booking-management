package repository

import (
	"booking-hotel/model/domain"
	"context"
	"database/sql"
)

type RoleRepository interface {
	Save(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role
	Update(ctx context.Context, tx *sql.Tx, role domain.Role) domain.Role
	Delete(ctx context.Context, tx *sql.Tx, role domain.Role)
	FindById(ctx context.Context, tx *sql.Tx, roleId int) (domain.Role, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Role
}
