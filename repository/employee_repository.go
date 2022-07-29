package repository // 2. relates to domain

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

//interface EmployeeRepository punya banyak method save-update-delete-findbyid-findall
/// bisa juga menggunakan struct, tapi tiap fucntion dari struct biasa nya berbeda lokasi

/*
what and who (definition and the user of the subject) :
- interface adalah kumpulan definisi method yg dibungkus dengan nama tertentu
-- method adalah fungsi yg menempel pada type dan memiliki akses ke property struct tertentu
-- bedanya dengan fungsi biasa, method punya pengaksesan dan deklarasi yg berbeda
-- method digunakan untuk  meng-enkapsulasi sebuah proses kerja

why and when (reason to use it, when you should and shouldn't use it) :

how and where (how to use, how it works) :

*/
