package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
	"time"
)

func UbahData(db *sql.DB, query string, user entities.User, before string) {
	var data string
	var option string
	// reader := bufio.NewReader(os.Stdin)
	fmt.Println("Data Sebelum : ", before)
	fmt.Print("Masukkan data baru : ")
	fmt.Scan(&data)
	// data, _ := reader.ReadString('\n')
	fmt.Println("Apakah anda yakin ingin mengubahnya?(y/n)")
	fmt.Scan(&option)
	// option, _ = reader.ReadString('\n')
	switch option {
	case "y", "Y":
		var datetime string = time.Now().Format("2006-01-02 15:04:05")

		_, err := db.Exec(query, data, datetime, user.Id)
		if err != nil {
			fmt.Println("Update data gagal", err.Error())
		} else {
			fmt.Println("Update data berhasil")
		}
	case "n", "N":
		fmt.Println("Batal menyimpan data")
	}
}

func UpdateAccount(db *sql.DB, user entities.User) {
	var menuupdateaccount = false
	for menuupdateaccount == false {
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
			menuupdateaccount = true
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
