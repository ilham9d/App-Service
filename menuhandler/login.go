package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
	"log"
)

func Login(db *sql.DB) {
	var user entities.User
	var telp, pass string

	fmt.Println("Nomor Telepon")
	fmt.Scanln(&telp)
	fmt.Println("Password")
	fmt.Scanln(&pass)

	row := db.QueryRow("select id, name, email, phone_number, password from accounts where phone_number = ?", telp)
	if err := row.Scan(&user.Id, &user.Nama, &user.Email, &user.PhoneNumber, &user.Password); err != nil {
		log.Fatal("error scan data", err.Error())
	}

	fmt.Println("phone:", user.PhoneNumber)
	fmt.Println("pass: ", user.Password)
	if telp == user.PhoneNumber && pass == user.Password {
		fmt.Println("login berhasil")
		MenuUtama() //memanggil menu utama
	} else {
		fmt.Println("no telpon atau password salah")
	}
}
