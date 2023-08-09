package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
)

func HistoryTopUp(db *sql.DB, user entities.User) {
	// var id, pengirim, penerima, balance, datetime string
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
	/*
		var pilihan int
		fmt.Println("0. Kembali")
		fmt.Print("Masukkan pilihan : ")
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 0:
			fmt.Println("")
		}
	*/
}
