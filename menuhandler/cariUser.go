package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func CariUser(db *sql.DB, user entities.User) {
	var phoneNumberToView string
	fmt.Println("Insert Phone_Number: ")
	fmt.Scanln(&phoneNumberToView)
	query := "SELECT name, email, phone_number FROM accounts WHERE phone_number = ?"
	row := db.QueryRow(query, phoneNumberToView)

	err := row.Scan(&user.Nama, &user.Email, &user.PhoneNumber)
	if err != nil {
		log.Fatal("error ", err.Error())
	}

	fmt.Println("Profile:")
	fmt.Println("Name:", user.Nama)
	fmt.Println("Email:", user.Email)
	fmt.Println("Phone_Number:", user.PhoneNumber)
}
