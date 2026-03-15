package main

import "fmt"

func main() {
	isBarangTersedia := true
	isBarangKosong := false

	fmt.Println("Apakah Barang Tersedia? :", isBarangTersedia)
	fmt.Println("Apakah Barang Kosong? :", isBarangKosong)

	fmt.Println()

	isInStock := true
	isAddressValid := false
	isPickupSelected := true
	isItemDamaged := false

	// Kita gunakan operator AND dan OR untuk mengecek apakah barang layak dikirim atau tidak.
	// lebih detail nya ada di file operasi-boolean/main.go
	isEligible := isInStock && (isAddressValid || isPickupSelected) && !isItemDamaged

	fmt.Println("---SISTEM VERIFIKASI GUDANG---")
	fmt.Printf("Status Stok          : %t\n", isInStock)
	fmt.Printf("Status Alamat        : %t\n", isAddressValid)
	fmt.Printf("Pickup di gudang     : %t\n", isPickupSelected)
	fmt.Printf("Apakah Barang Rusak? : %t\n", isItemDamaged)
	fmt.Println("--------------------------------")
	fmt.Printf("Apakah Barang Layak Dikirim? : %t\n", isEligible)

	fmt.Println()

	// Kita gunakan IF STATEMENT untuk mengecek apakah barang layak dikirim atau tidak.
	// lebih detail nya ada di file if-statement/main.go
	if isEligible {
		fmt.Println("Silahkan Cetak Label Pengiriman.")
	} else {
		fmt.Println("Barang tidak laya di kirim!")
	}

}

//
// CATATAN TIPE DATA BOOLEAN & FORMAT PRINT
//
// Tipe data boolean:
//   - Hanya dua nilai: true dan false. Dipakai untuk kondisi, flag, validasi.
//   - Nama variabel sering diawali is/has/can (isBarangTersedia, isEligible) agar mudah dibaca.
//
// Format cetak (Printf):
//   - %t : untuk boolean. Menampilkan "true" atau "false". Pakai Printf, bukan Println, agar %t diolah.
//   - Println("teks", x) : mencetak teks lalu nilai x (tanpa format verb). Cocok untuk output sederhana.
//
// Operator boolean (lihat operasi-boolean/main.go untuk detail):
//   - && : AND (kedua kondisi harus true).
//   - || : OR (salah satu true cukup).
//   - !  : NOT (membalik nilai).
//
