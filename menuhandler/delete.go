package menuhandler

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func SoftDelete(db *sql.DB, id string) error {
	currentTime := time.Now()
	fmt.Println(currentTime)
	_, err := db.Exec("UPDATE accounts SET delete_at = ? WHERE id = ?", currentTime, id)
	if err != nil {
		return err
	}
	return nil
}

func Delete() {
	var db *sql.DB
	var err error
	var connectionString = os.Getenv("CONNECTION_DB")

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("error connection", err.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("error ping db", err.Error())
	} else {
		fmt.Println("success ping")
	}

	defer db.Close()

	var idToDelete string
	fmt.Print("Insert your ID: ")
	fmt.Scanln(&idToDelete)

	err = SoftDelete(db, idToDelete)
	if err != nil {
		log.Fatal("Failed to Delete", err.Error())
	} else {
		fmt.Println("Delete Successfully!")
	}
}
