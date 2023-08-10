package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
)

func MenuUtama(db *sql.DB, user entities.User) {

	fmt.Println("Menu Utama")
	var menuUtama int
	fmt.Println("1. Lihat Saldo\n2. Top-up\n3. Transfer\n4. Riwayat Transaksi\n5. Lihat Profil\n6. Hapus Akun\n7. Cari User\n8. Keluar Dari Sistem")
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
		fmt.Println("Keluar")
	}
}
