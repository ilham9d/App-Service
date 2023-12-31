package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
)

func getLoginData(db *sql.DB, id string) entities.User {
	var user entities.User
	var create, update sql.NullString
	row := db.QueryRow("select id, name, email, phone_number, password, balance, create_at, update_at from accounts where id = ? and delete_at is null", id)
	err := row.Scan(&user.Id, &user.Nama, &user.Email, &user.PhoneNumber, &user.Password, &user.Balance, &create, &update)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if create.Valid {
			user.Create_at = create.String
		}
		if update.Valid {
			user.Update_at = update.String
		}
	}
	return user
}

func MenuUtama(db *sql.DB, id string) {

	var utama = false
	for !utama {
		var user = getLoginData(db, id)
		fmt.Println("Menu Utama")
		var menuUtama int
		fmt.Println("1. Lihat Saldo\n2. Top-up\n3. Transfer\n4. Riwayat Transaksi\n5. Lihat Profil\n6. Hapus Akun\n7. Cari User\n8. Log Out")
		fmt.Print("Masukkan Pilihan : ")
		fmt.Scanln(&menuUtama)
		switch menuUtama {
		case 1:
			fmt.Println("==================================")
			ShowSaldo(db, user)
		case 2:
			fmt.Println("==================================")
			TopUp(db, user)
		case 3:
			fmt.Println("==================================")
			Transfer(db, user)
		case 4:
			fmt.Println("==================================")
			History(db, user)
		case 5:
			fmt.Println("==================================")
			ReadProfile(db, user)
		case 6:
			fmt.Println("==================================")
			Delete(db, user)
		case 7:
			fmt.Println("==================================")
			CariUser(db, user)
		case 8:
			fmt.Println("==================================")
			fmt.Println("Anda keluar dari akun anda")
			fmt.Println("==================================")
			utama = true
		default:
			fmt.Println("Pilihan tidak valid")
			fmt.Println("==================================")
		}
	}
}
