package main

import "fmt"

func main() {
	// Variabel string (literal dalam tanda petik ganda)
	namaGudang := "Gudang Milik KIM"
	kodeBarang := "001_KIM_Stuff"

	// len() mengembalikan jumlah byte (bukan jumlah karakter; untuk UTF-8 bisa lebih kompleks)
	jumlahKarakter := len(kodeBarang)

	// Indeks string menghasilkan byte; cast ke string() untuk dapat karakter
	hurufPertama := string(kodeBarang[0])
	hurufKedua := string(kodeBarang[1])
	hurufKetiga := string(kodeBarang[2])
	ambilKode := string(kodeBarang[4:7])

	// Penggabungan string dengan operator +
	kategori := hurufPertama + hurufKedua + hurufKetiga

	fmt.Println("---SISTEM LABELING BARANG---")
	fmt.Println("Nama Gudang : ", namaGudang)
	fmt.Println("Kode Barang : ", kodeBarang)
	fmt.Println("Panjang Kode: ", jumlahKarakter)
	fmt.Println("Singkat Tipe: ", kategori)
	fmt.Println("Kode Alias: ", ambilKode)

	// Perbandingan string dengan ==
	// lebih lanjut tentang perbandingan string ada di file operasi-string/main.go
	if kategori == "001" {
		fmt.Println("Ini adalah barang pertama!")
	} else {
		fmt.Println("Ini bukan barang yang pertama!")
	}
}

//
// CATATAN TIPE DATA STRING
//
// Tipe data string:
//   - Deretan byte (immutable). Literal dengan tanda petik ganda "..." atau backtick `...`.
//   - Backtick: string mentah (raw), newline dan escape tidak diolah. Petik ganda: \n, \t, dll. diolah.
//
// Operasi umum:
//   - len(s)       : panjang string dalam byte. Untuk teks ASCII 1 byte = 1 karakter; UTF-8 bisa multi-byte per karakter.
//   - s[i]         : byte ke-i (tipe byte/uint8). Harus di-cast: string(s[i]) untuk tampil sebagai karakter.
//   - s[low:high]  : slice string dari indeks low sampai sebelum high (high tidak ikut). Contoh: "001_KIM_Stuff"[4:7] → "KIM".
//   - s + t        : penggabungan string (konkatenasi).
//   - ==, !=, <, > : perbandingan string (leksikografis, berdasarkan byte).
//
// Lihat juga: format-verb (Printf %s), slice-array, operasi string di package strings (strings.Split, strings.Contains, dll).
//
