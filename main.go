package main

import (
	"app-service/menuhandler"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
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

	var pilihan int
	fmt.Println("Selamat datang")
	fmt.Println("pilih menu")
	fmt.Print("1. Login\n2. Register\n3. About\n")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scanln(&pilihan)

	for true {
		switch pilihan {
		case 1:
			menuhandler.Login()
		case 2:
			fmt.Println("register")
		case 3:
			fmt.Println("About")
		}
	}
}
