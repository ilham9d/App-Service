package menuhandler

import (
	"database/sql"
	"fmt"
)

func Login(db *sql.DB) {
	var id, phoneNumber, password string
	var telp, pass string

	fmt.Print("Nomor Telepon : ")
	fmt.Scanln(&telp)
	fmt.Print("Password : ")
	fmt.Scanln(&pass)

	row := db.QueryRow("select id, phone_number, password from accounts where phone_number = ? and delete_at is null", telp)
	if err := row.Scan(&id, &phoneNumber, &password); err != nil {
		fmt.Println("Akun tidak terdaftar ")
	} else {
		if telp == phoneNumber && pass == password {
			fmt.Println("login berhasil!")
			MenuUtama(db, id) //memanggil menu utama
		} else {
			fmt.Println("no telpon atau password salah.")
		}
	}
}
