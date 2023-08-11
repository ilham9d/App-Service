package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
	"log"
)

func ShowSaldo(db *sql.DB, user entities.User) {
	var menusaldo = false

	for !menusaldo {
		var pilihan int
		var saldo int
		row := db.QueryRow("select balance from accounts where id = ? and delete_at is null", user.Id)
		if err := row.Scan(&saldo); err != nil {
			log.Fatal("error scan data ", err.Error())
		}
		fmt.Println("Saldo Anda Saat Ini :")
		fmt.Println("Rp. ", saldo)
		fmt.Println()
		fmt.Println("1. Kembali")
		fmt.Print("Masukkan pilihan : ")
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			fmt.Println("==================================")
			menusaldo = true
		default:
			fmt.Println("Pilihan tidak valid")
			fmt.Println("==================================")
		}
	}
}
