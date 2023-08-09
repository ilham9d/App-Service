package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
	"log"
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

func Delete(db *sql.DB, user entities.User) {
	fmt.Print("Are you sure you want to delete this user? (yes/no): ")
	var confirmation string
	_, err := fmt.Scanln(&confirmation)
	if err != nil {
		log.Fatal("Failed to read input:", err.Error())
	}

	if confirmation == "yes" {
		err := SoftDelete(db, user.Id)
		if err != nil {
			log.Fatal("Failed to Delete", err.Error())
		} else {
			fmt.Println("Delete Successfully!")
		}
	} else {
		fmt.Println("Delete operation canceled.")
	}
}
