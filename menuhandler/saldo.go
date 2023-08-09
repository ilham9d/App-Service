package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
	"log"
)

func ShowSaldo(db *sql.DB, user entities.User) {
	// var pilihan int

	row := db.QueryRow("select balance from accounts where id = ? and delete_at is null", user.Id)
	if err := row.Scan(&user.Balance); err != nil {
		log.Fatal("error scan data ", err.Error())
	}
	fmt.Println("Saldo Anda Saat Ini :")
	fmt.Println("Rp. ", user.Balance)
	// fmt.Println()
	// fmt.Println("1. Kembali\n2. Keluar dari sistem")
	// fmt.Print("Masukkan pilihan : ")
	// fmt.Scan(&pilihan)
	// switch pilihan {
	// case 1:

	// }
}
