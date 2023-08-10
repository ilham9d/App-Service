package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
)

func getLoginData(db *sql.DB, id string) (entities.User, bool) {
	var flag bool
	var user entities.User
	var create, update sql.NullString
	row := db.QueryRow("select id, name, email, phone_number, password, balance, create_at, update_at from accounts where id = ? and delete_at is null", id)
	err := row.Scan(&user.Id, &user.Nama, &user.Email, &user.PhoneNumber, &user.Password, &user.Balance, &create, &update)
	if err != nil {
		fmt.Println(err.Error())
		flag = true
	} else {
		if create.Valid {
			user.Create_at = create.String
		}
		if update.Valid {
			user.Update_at = update.String
		}
		flag = false
	}
	return user, flag
}

func MenuUtama(db *sql.DB, id string) {

	var utama = false
	for utama == false {
		var user, flag = getLoginData(db, id)
		utama = flag
		fmt.Println("Menu Utama")
		var menuUtama int
		fmt.Println("1. Lihat Saldo\n2. Top-up\n3. Transfer\n4. Riwayat Transaksi\n5. Lihat Profil\n6. Hapus Akun\n7. Cari User\n8. Log Out")
		fmt.Print("Masukkan Pilihan : ")
		fmt.Scanln(&menuUtama)
		switch menuUtama {
		case 1:
			ShowSaldo(db, user)
		case 2:
			TopUp(db, user)
		case 3:
			Transfer(db, user)
		case 4:
			History(db, user)
		case 5:
			ReadProfile(db, user)
		case 6:
			Delete(db, user)
		case 7:
			CariUser(db, user)
		case 8:
			utama = true
		}
	}
	fmt.Println("Anda keluar dari akun anda")
}
