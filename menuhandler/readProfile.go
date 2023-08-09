package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
	"log"
)

func ReadProfile(db *sql.DB, user entities.User) {
	var create, update sql.NullString

	row := db.QueryRow("select id, name, email, phone_number, password, balance, create_at, update_at from accounts where id = ? and delete_at is null", user.Id)
	if err := row.Scan(&user.Id, &user.Nama, &user.Email, &user.PhoneNumber, &user.Password, &user.Balance, &create, &update); err != nil {
		log.Fatal("error scan data ", err.Error())
	}
	if create.Valid {
		user.Create_at = create.String
	}
	if update.Valid {
		user.Update_at = update.String
	}

	fmt.Print("Nama : ", user.Nama, "\nEmail : ", user.Email, "\nNomor Telepon : ", user.PhoneNumber, "\nDibuat pada : ", user.Create_at, "\n")
	if user.Update_at != "" {
		fmt.Print("Diupdate pada : ", user.Update_at, "\n")
	}

	var pilihan int
	fmt.Println("1. Ubah Profil")
	fmt.Println("2. Kembali")
	fmt.Print("Masukkan Pilihan : ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		UpdateAccount(db, user)
	case 2:
		MenuUtama(db, user)
	}
}
