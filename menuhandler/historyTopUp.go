package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
)

func HistoryTopUp(db *sql.DB, user entities.User) {

	var menuhistorytopup = false
	for !menuhistorytopup {
		rows, err := db.Query("select id, account_id_pengirim, account_id_penerima, status, balance, date_time_transaction from transaction where status = ? and account_id_penerima = ? and account_id_pengirim = ? order by date_time_transaction desc", "Top-Up", user.Id, user.Id)
		if err != nil {
			fmt.Println("query error ", err.Error())
		}

		var transaction []entities.Transaction
		for rows.Next() {
			var row entities.Transaction
			errScan := rows.Scan(&row.Id, &row.SenderId, &row.ReceiverId, &row.Status, &row.Balance, &row.Time)
			if errScan != nil {
				fmt.Println("Scan error ", errScan.Error())
			} else {
				transaction = append(transaction, row)
			}
		}
		for _, v := range transaction {
			fmt.Print(v.Status, "  ", v.Balance, "  ", v.Time, "\n")
		}
		fmt.Println("------Akhir daftar Transaksi------")

		var pilihan int
		fmt.Println("1. Kembali")
		fmt.Print("Masukkan pilihan : ")
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			fmt.Println("==================================")
			menuhistorytopup = true
		default:
			fmt.Println("Pilihan tidak valid")
			fmt.Println("==================================")
		}
	}
}
