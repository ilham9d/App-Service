package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
)

func UbahData(db *sql.DB, query string, user entities.User, before string) {
	var data string
	var option string
	fmt.Println("Data Sebelum : ", before)
	fmt.Print("Masukkan data baru : ")
	fmt.Scan(&data)
	fmt.Println("Apakah anda yakin ingin mengubahnya?(y/n)")
	fmt.Scan(&option)
	switch option {
	case "y", "Y":
		_, err := db.Exec(query, data, user.Id)
		if err != nil {
			fmt.Println("Update data gagal", err.Error())
		} else {
			fmt.Println("Update data berhasil")
		}
	case "n", "N":
		fmt.Println("Batal menyimpan data")
	}
}

func UpdateAccount(db *sql.DB, user entities.User) {
	var pilihan int
	fmt.Println("Pilih data yang akan diubah")
	fmt.Println("1. Nama\n2. Email\n3. Nomor Telepon\n4. Password")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		query := "Update accounts set name = ? where id = ?"
		before := user.Nama
		UbahData(db, query, user, before)
	case 2:
		query := "Update accounts set email = ? where id = ?"
		before := user.Email
		UbahData(db, query, user, before)
	case 3:
		query := "Update accounts set phone_number = ? where id = ?"
		before := user.PhoneNumber
		UbahData(db, query, user, before)
	case 4:
		query := "Update accounts set password = ? where id = ?"
		before := user.Password
		UbahData(db, query, user, before)
	default:
	}
}
