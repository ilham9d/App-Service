package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
)

func History(db *sql.DB, user entities.User) {
	var pilihan int
	fmt.Println("1. Riwayat Top-Up")
	fmt.Println("2. Riwayat Transfer")
	fmt.Print("Masukkan pilihan : ")
	fmt.Scanln(&pilihan)
	switch pilihan {
	case 1:
		HistoryTopUp(db, user)
	case 2:
		fmt.Println("Transfer")
	case 3:
		fmt.Println("Kembali")
	}
}
