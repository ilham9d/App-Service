package menuhandler

import (
	"app-service/entities"
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"
)

func UbahData(db *sql.DB, query string, user entities.User, before string) {
	var data string
	var option string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Data Sebelum : ", before)
	fmt.Print("Masukkan data baru : ")
	data, _ = reader.ReadString('\n')
	data = strings.TrimSpace(data)

	fmt.Print("Apakah anda yakin ingin mengubahnya?(y/n)")
	fmt.Scanln(&option)
	switch option {
	case "y", "Y":
		var datetime string = time.Now().Format("2006-01-02 15:04:05")

		_, err := db.Exec(query, data, datetime, user.Id)
		if err != nil {
			fmt.Println("Update data gagal", err.Error())
			fmt.Println("==================================")
		} else {
			fmt.Println("Update data berhasil")
			fmt.Println("==================================")
		}
	case "n", "N":
		fmt.Println("Batal menyimpan data")
		fmt.Println("==================================")
	default:
		fmt.Println("Pilihan tidak valid, batal menyimpan data")
		fmt.Println("==================================")
	}
}

func UpdateAccount(db *sql.DB, user entities.User) {
	var menuupdateaccount = false
	for !menuupdateaccount {
		var pilihan int
		fmt.Println("Pilih data yang akan diubah")
		fmt.Println("1. Nama\n2. Email\n3. Nomor Telepon\n4. Password\n5. Kembali")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			query := "Update accounts set name = ?, update_at = ? where id = ?"
			before := user.Nama
			UbahData(db, query, user, before)
		case 2:
			query := "Update accounts set email = ?, update_at = ? where id = ?"
			before := user.Email
			UbahData(db, query, user, before)
		case 3:
			query := "Update accounts set phone_number = ?, update_at = ? where id = ?"
			before := user.PhoneNumber
			UbahData(db, query, user, before)
		case 4:
			query := "Update accounts set password = ?, update_at = ? where id = ?"
			before := user.Password
			UbahData(db, query, user, before)
		case 5:
			fmt.Println("==================================")
			menuupdateaccount = true
		default:
			fmt.Println("Pilihan tidak valid")
			fmt.Println("==================================")
		}
	}
}
