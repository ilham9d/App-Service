package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
)

func ReadProfile(db *sql.DB, user entities.User) {

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
