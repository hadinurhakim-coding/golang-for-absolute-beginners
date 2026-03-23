package main

import "fmt"

func main() {
	// defer menunda eksekusi sampai fungsi pemanggil selesai (return atau panic).
	// Urutan: LIFO — yang di-defer terakhir dijalankan pertama.
	fmt.Println("mulai")
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	fmt.Println("tengah")
	// Output urutan cetak: mulai, tengah, defer 3, defer 2, defer 1

	contohCleanup()

	// Argumen defer dievaluasi saat baris defer dijalankan, bukan saat fungsi defer dipanggil
	i := 0
	defer fmt.Println("defer i (nilai saat defer):", i)
	i = 10
	fmt.Println("i setelah ubah:", i)
}

func contohCleanup() {
	fmt.Println("--- contoh pola cleanup ---")
	defer fmt.Println("tutup sumber daya (simulasi)")
	fmt.Println("kerja dengan sumber daya")
	// Pola nyata: defer file.Close(), mutex.Unlock(), dll.
}
