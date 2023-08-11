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

func Delete(db *sql.DB, user entities.User) { //db *sql.DB, user entities.User

	// var db *sql.DB
	// var err error
	// var connectionString = os.Getenv("CONNECTION_DB")

	// db, err = sql.Open("mysql", connectionString)
	// if err != nil {
	// 	log.Fatal("error connection", err.Error())
	// }

	// errPing := db.Ping()
	// if errPing != nil {
	// 	log.Fatal("error ping db", err.Error())
	// } else {
	// 	fmt.Println("success ping")
	// }

	// defer db.Close()

	// var idToDelete string
	// fmt.Print("Insert your ID: ")
	// fmt.Scanln(&idToDelete)
	var pilihan string
	fmt.Println("Apakah anda yakin ingin menghapus akun anda?(y/n)")
	fmt.Print("Masukkan pilihan : ")
	fmt.Scanln(&pilihan)
	switch pilihan {
	case "y", "Y":
		err := SoftDelete(db, user.Id)
		if err != nil {
			log.Fatal("Failed to Delete", err.Error())
		} else {
			fmt.Println("Delete Successfully!")
		}
	case "n", "N":
		fmt.Println("Batal menghapus akun")
	default:
		fmt.Println("Pilihan tidak valid")
	}
}
