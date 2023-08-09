package menuhandler

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func Register(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)

	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	newid, err := gonanoid.Generate(str, 20)
	if err != nil {
		log.Default()
	}

	fmt.Print("Masukkan nama anda: ")
	Name, _ := reader.ReadString('\n')
	Name = strings.TrimSpace(Name)

	fmt.Print("Masukkan password anda: ")
	Password, _ := reader.ReadString('\n')
	Password = strings.TrimSpace(Password)

	fmt.Print("Masukkan email anda: ")
	Email, _ := reader.ReadString('\n')
	Email = strings.TrimSpace(Email)

	fmt.Print("Masukkan nomor telpon anda :")
	Phone_Number, _ := reader.ReadString('\n')
	Phone_Number = strings.TrimSpace(Phone_Number)

	var option string
	fmt.Println("Apakah data anda sudah benar?(y/n)")
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scanln(&option)
	switch option {
	case "y", "Y":
		result, errInsert := db.Exec("INSERT INTO accounts (id, name, password, email, phone_number) VALUES (?, ?, ?, ?, ?)", newid, Name, Password, Email, Phone_Number)
		if errInsert != nil {
			log.Fatal("error insert", errInsert.Error())
		} else {
			row, _ := result.RowsAffected()
			if row > 0 {
				fmt.Println("Berhasil mendaftar")
			} else {
				fmt.Println("Gagal mendaftar")
			}
		}
	case "n", "N":
		fmt.Println("Batal mendaftar")
	}

}
