package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
	"log"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func Transfer(db *sql.DB, user entities.User) {
	var amount int
	var option string
	var inputPhoneNumber string

	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	newid, err := gonanoid.Generate(str, 20)
	if err != nil {
		log.Default()
	}

	fmt.Println("Tranfer:")
	fmt.Println("1. Pilih nomor tujuan yang pernah digunakan")
	fmt.Println("2. Masukkan Nomor Penerima baru")
	fmt.Print("Pilihan Anda: ")
	fmt.Scanln(&option)

	switch option {
	case "1":
		fmt.Println("Nomor tujuan yang pernah digunakan:")

		fmt.Print("Pilih nomor tujuan: ")
		fmt.Scanln(&inputPhoneNumber)

	case "2":
		fmt.Print("Masukkan Nomor Penerima baru: ")
		fmt.Scanln(&inputPhoneNumber)
	default:
		fmt.Println("Opsi tidak valid")
		return
	}

	var recipientId string
	var recipientName string
	var recipientEmail string
	var recipientPhoneNumber string
	var recipientBalance int

	rowrecipient := db.QueryRow("SELECT id, balance, name, phone_number, email FROM accounts WHERE phone_number = ? AND delete_at IS null", inputPhoneNumber)
	if errScanRec := rowrecipient.Scan(&recipientId, &recipientBalance, &recipientName, &recipientPhoneNumber, &recipientEmail); err != nil {
		log.Fatal("error scanning recipient's balance: ", errScanRec.Error())
	}

	fmt.Println("Tujuan anda")
	fmt.Print("Nama : ", recipientName, "\n")
	fmt.Print("Email : ", recipientEmail, "\n")
	fmt.Print("No. Telpon : ", recipientPhoneNumber, "\n")

	fmt.Print("Masukkan Nominal Transfer : ")
	fmt.Print("Rp. ")
	fmt.Scanln(&amount)

	fmt.Print("Apakah anda ingin melakukan transfer? (y/n)")
	fmt.Scanln(&option)

	switch option {
	case "y", "Y":
		var senderBalance int
		row := db.QueryRow("SELECT balance FROM accounts WHERE id = ? AND delete_at IS null", user.Id)
		if err := row.Scan(&senderBalance); err != nil {
			log.Fatal("error scanning sender's balance: ", err.Error())
		}
		if senderBalance < amount {
			fmt.Println("Saldo tidak mencukupi untuk melakukan transfer sebesar Rp.", amount)
			return
		}

		// Update sender's balance
		newSenderBalance := senderBalance - amount
		_, err := db.Exec("UPDATE accounts SET balance = ? WHERE id = ?", newSenderBalance, user.Id)
		if err != nil {
			fmt.Println("Transfer gagal ", err.Error())
			return
		}

		// Update recipient's balance
		newRecipientBalance := recipientBalance + amount
		_, err = db.Exec("UPDATE accounts SET balance = ? WHERE id = ?", newRecipientBalance, recipientId)
		if err != nil {
			fmt.Println("Transfer gagal ", err.Error())
			return
		}
		var datetime string = time.Now().Format("2006-01-02 15:04:05")
		_, err = db.Exec("INSERT INTO transaction (id, account_id_pengirim, account_id_penerima, status, balance, date_time_transaction) VALUES (?,?,?,?,?,?)", newid, user.Id, recipientId, "Transfer", amount, datetime)
		if err != nil {
			fmt.Println("Transfer gagal ", err.Error())
			return
		}
		fmt.Println("Transfer berhasil")
		fmt.Println("==================================")
	case "n", "N":
		fmt.Println("Transfer dibatalkan")
		fmt.Println("==================================")
	default:
		fmt.Println("Pilihat tidak valid, Transfer dibatalkan")
		fmt.Println("==================================")
	}
}
