package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
)

func HistoryTransfer(db *sql.DB, user entities.User) {
	var menuhistorytransfer = false
	for menuhistorytransfer == false {
		rows, err := db.Query("select pengirim.name as nama_pengirim, penerima.name as nama_penerima, pengirim.phone_number, penerima.phone_number, t.balance, t.date_time_transaction from transaction as t inner join accounts as pengirim on t.account_id_pengirim = pengirim.id inner join accounts as penerima on t.account_id_penerima = penerima.id where (pengirim.id = ? or penerima.id=?) and status = ?;", user.Id, user.Id, "Transfer")
		if err != nil {
			fmt.Println("query error ", err.Error())
		}

		var transfer []entities.Transfer
		for rows.Next() {
			var row entities.Transfer
			errScan := rows.Scan(&row.Pengirim, &row.Penerima, &row.PhonePengirim, &row.PhonePenerima, &row.Balance, &row.Time)
			if errScan != nil {
				fmt.Println("Scan error ", errScan.Error())
			} else {
				transfer = append(transfer, row)
			}
		}

		for _, v := range transfer {
			var status string
			if v.Pengirim == user.Nama {
				status = "Keluar"
				fmt.Print(status, " ", v.Penerima, " ", v.PhonePenerima, " ", v.Balance, " ", v.Time, "\n")
			} else if v.Penerima == user.Nama {
				status = "Masuk"
				fmt.Print(status, " ", v.Pengirim, " ", v.PhonePengirim, " ", v.Balance, " ", v.Time, "\n")
			}
		}
		fmt.Println("------Akhir daftar transaksi------")
		var pilihan int
		fmt.Println("1. Kembali")
		fmt.Print("Masukkan pilihan : ")
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			menuhistorytransfer = true
		}
	}
}
