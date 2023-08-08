package menuhandler

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

func Register(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Insert Your ID: ")
	ID, _ := reader.ReadString('\n')
	ID = strings.TrimSpace(ID)

	fmt.Print("Insert Your Name: ")
	Name, _ := reader.ReadString('\n')
	Name = strings.TrimSpace(Name)

	fmt.Print("Insert Your Password: ")
	Password, _ := reader.ReadString('\n')
	Password = strings.TrimSpace(Password)

	fmt.Print("Insert Your Email: ")
	Email, _ := reader.ReadString('\n')
	Email = strings.TrimSpace(Email)

	fmt.Print("Insert Your Phone Number: ")
	Phone_Number, _ := reader.ReadString('\n')
	Phone_Number = strings.TrimSpace(Phone_Number)

	result, errInsert := db.Exec("INSERT INTO accounts (id, name, password, email, phone_number) VALUES (?, ?, ?, ?, ?)", ID, Name, Password, Email, Phone_Number)
	if errInsert != nil {
		log.Fatal("error insert", errInsert.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("success to register")
		} else {
			fmt.Println("failed to register")
		}
	}
}
