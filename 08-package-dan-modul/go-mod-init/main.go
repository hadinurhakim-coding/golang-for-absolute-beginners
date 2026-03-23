package main

import "fmt"

func main() {
	// Modul repo ini sudah diinisialisasi dengan: go mod init golang-for-absolute-beginners
	// Perintah umum lain:
	//   go mod tidy   — rapikan dependensi di go.mod / go.sum
	//   go get pkg@versi — tambah dependensi
	//   go list -m all — daftar modul
	fmt.Println("Lihat go.mod di root repo untuk nama modul dan versi Go.")
}
