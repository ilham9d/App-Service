package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
	"log"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func TopUp(db *sql.DB, user entities.User) {
	var nominal int
	var option string

	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	newid, err := gonanoid.Generate(str, 20)
	if err != nil {
		log.Default()
	}

	fmt.Println("Masukkan Nominal Top-Up : ")
	fmt.Print("Rp. ")
	fmt.Scanln(&nominal)
	fmt.Print("Apakah anda ingin melakukan Top-Up sebesar ", nominal, "?(y/n)")
	fmt.Scanln(&option)

	switch option {
	case "y", "Y":
		row := db.QueryRow("select balance from accounts where id = ? and delete_at is null", user.Id)
		if err := row.Scan(&user.Balance); err != nil {
			log.Fatal("error scan data ", err.Error())
		}

		var newNominal = user.Balance + nominal

		result, err := db.Exec("Update accounts set balance = ? where id = ?", newNominal, user.Id)
		if err != nil {
			fmt.Println("Top-Up gagal ", err.Error())
		} else {
			row, _ := result.RowsAffected()
			if row > 1 {
				log.Fatal(err.Error())
			}
			var datetime string = time.Now().Format("2006-01-02 15:04:05")
			_, err2 := db.Exec("insert into transaction(id, account_id_pengirim, account_id_penerima, status, balance, date_time_transaction) values(?,?,?,?,?,?)", newid, user.Id, user.Id, "Top-Up", nominal, datetime)
			if err2 != nil {
				log.Fatal("error insert data transaction ", err2.Error())
			}
			fmt.Println("Top-Up berhasil")
			fmt.Println("==================================")
		}
	case "n", "N":
		fmt.Println("Top-Up dibatalkan")
		fmt.Println("==================================")
	default:
		fmt.Println("Pilihan tidak valid, Top-Up dibatalkan")
		fmt.Println("==================================")
	}
}
