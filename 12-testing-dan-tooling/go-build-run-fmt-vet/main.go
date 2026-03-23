package main

import "fmt"

func main() {
	// Tooling umum (dari root modul atau folder ini):
	//   go build .     — kompilasi biner
	//   go run .       — kompilasi + jalankan
	//   go fmt ./...   — format semua paket di bawah modul
	//   go vet ./...   — analisis statis (curiga bug)
	//   go mod tidy    — bersihkan dependensi
	fmt.Println("Baca komentar di main.go untuk daftar perintah.")
}
