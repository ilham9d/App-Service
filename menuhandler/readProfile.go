package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
)

func ReadProfile(db *sql.DB, user entities.User) {

	fmt.Print("Nama : ", user.Nama, "\nEmail : ", user.Email, "\nNomor Telepon : ", user.PhoneNumber, "\nDibuat pada : ", user.Create_at)
	if user.Update_at != "" {
		fmt.Println("Diupdate pada : ", user.Update_at)
	}
}
