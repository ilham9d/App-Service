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
	var create, update, delete sql.NullString

	fmt.Println("Nomor Telepon")
	fmt.Scan(&telp)
	fmt.Println("Password")
	fmt.Scan(&pass)

	row := db.QueryRow("select id, name, email, phone_number, password, balance, create_at, update_at, delete_at from accounts where phone_number = ? and delete_at is null", telp)
	if err := row.Scan(&user.Id, &user.Nama, &user.Email, &user.PhoneNumber, &user.Password, &user.Balance, &create, &update, &delete); err != nil {
		log.Fatal("error scan data ", err.Error())
	}
	if create.Valid {
		user.Create_at = create.String
	}
	if update.Valid {
		user.Update_at = update.String
	}
	if delete.Valid {
		user.Delete_at = delete.String
	}

	fmt.Println("phone:", user.PhoneNumber)
	fmt.Println("pass: ", user.Password)
	if telp == user.PhoneNumber && pass == user.Password {
		fmt.Println("login berhasil")
		MenuUtama(db, user) //memanggil menu utama
	} else {
		fmt.Println("no telpon atau password salah")
	}
}
