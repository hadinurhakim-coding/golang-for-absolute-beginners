// Program pertama: menampilkan teks "Hello, everybody!" di layar

/*
	package main punya peran khusus di Go:

	1. Menandakan ini adalah program yang bisa dijalankan (executable), bukan library.
	2. Hanya paket bernama "main" yang bisa di-build/di-run menjadi program berdiri sendiri.
	3. Di dalam package main harus ada fungsi main() — itu yang pertama kali dipanggil saat program jalan.
	4. Package lain (misalnya "utils", "model") dipakai untuk kode yang di-import oleh program; mereka tidak punya main() dan tidak bisa dijalankan langsung.

	Singkatnya: package main = "ini program utamaku", dan Go mencari func main() di sini sebagai titik mulai.
*/
package main


import "fmt" // Paket untuk mencetak teks ke layar (fmt = format)

func main() {
	// main() adalah titik awal program; Go akan menjalankan kode di sini terlebih dahulu
	fmt.Println("Hello, everybody!") // Cetak teks ke layar lalu pindah baris baru
}