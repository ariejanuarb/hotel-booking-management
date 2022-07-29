package domain //1. relates to repository

type Employee struct {
	Id       int
	Name     string
	Gender   string
	Email    string
	Password string
	RoleId   int
}

/*
what and who (definition and the user of the subject) :
- struct adalah kumpulan dari field yg dideklarasikan kedalam sebuah tipe data baru
- struct adalah penyederhanaan konsep class pada OOP, struct = class-versi-ringan
- struct tidak mendukung konsep inheritance (OOP) namun mendukung konsep komposisi = pemanggilan bersamaan dari banyak struct
- bedanya dengan interface, struct adalah kumpulan definisi variabel yg dibungkus sebagai tipe data baru


why and when (reason to use it, when you should and shouldn't use it) :
- sturct bisa memungkinkan kita untuk memberikan data struktur yg kompleks ke dalam sebuah sistem
- ketika kita harus merepresentasikan real-world entity yg memiliki banyak properties/fields
- ketika kita membutuhkan object baru yg berisi dari gabungan banyak tipe data
- kita bisa membuat sebuah variabel dan function berdasarkan object struct yg dibuat sebagai tipe datanya

how and where (how to use, how it works) :
- how to acces its fields = by using dot "." operatpr
- In composition, base structs can be embedded into a child struct
- and the methods of the base struct can be directly called on the child struct

*/
