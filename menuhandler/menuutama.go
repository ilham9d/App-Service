package menuhandler

import (
	"app-service/entities"
	"database/sql"
	"fmt"
)

func MenuUtama(db *sql.DB, user entities.User) {

	fmt.Println("Menu Utama")
	var menuUtama int
	fmt.Println("1. Lihat Saldo\n2. Top-up\n3. Transfer\n4. Riwayat Transaksi\n5. Lihat Profil\n6. Cari User\n7. Keluar Dari Sistem")
	fmt.Println("Masukkan Pilihan")
	fmt.Scan(&menuUtama)
	switch menuUtama {
	case 1:
		fmt.Println("lihat saldo")
		ShowSaldo()
	case 2:
		fmt.Println("top-Up")
	case 3:
		fmt.Println("Transfer")
	case 4:
		fmt.Println("Riwayat transaksi")
	case 5:
		ReadProfile(db, user)
	case 6:
		fmt.Println("Cari User")
	case 7:
		fmt.Println("Keluar")
	}
}
