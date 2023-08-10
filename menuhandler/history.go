package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
)

func History(db *sql.DB, user entities.User) {

	var menuhistory = false

	for menuhistory == false {
		var pilihan int
		fmt.Println("1. Riwayat Top-Up")
		fmt.Println("2. Riwayat Transfer")
		fmt.Println("3. Kembali")
		fmt.Print("Masukkan pilihan : ")
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			HistoryTopUp(db, user)
		case 2:
			HistoryTransfer(db, user)
		case 3:
			menuhistory = true
		}
	}
}
