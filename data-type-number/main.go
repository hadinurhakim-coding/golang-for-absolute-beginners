package main

import "fmt"

func main() {
	// 1. Menggunakan type inference (:=) agar lebih ringkas
	// Go akan otomatis mendeteksi tipe dari nilai yang diberikan
	umur := 25
	suhu := int8(-10)          // Casting diperlukan jika ingin tipe spesifik (int8)
	rataRataUmur := uint(60)   // camelCase untuk nama variabel multi-kata
	hargaSatuDollar := 16000.0 // Literal .0 → float64 (default untuk desimal)
	pi := float32(3.14)

	fmt.Printf("Umur : %d\n", umur)
	fmt.Printf("Suhu : %d\n", suhu)
	fmt.Printf("Rata-rata Umur : %d\n", rataRataUmur)
	fmt.Printf("Harga Satu Dollar : Rp.%.0f\n", hargaSatuDollar)
	fmt.Printf("Nilai pi : %.2f\n", pi)

	// Pindah baris dengan cara yang lebih simpel
	fmt.Println()

	// 2. Contoh kasus transaksi gudang — variabel terkait dikelompokkan dalam satu blok
	var (
		idBarang    uint32  = 90210
		stok        int     = 100
		hargaSatuan float64 = 15000.64
		jumlahBeli  int     = 5
	)

	totalBayar := hargaSatuan * float64(jumlahBeli)
	stokBaru := stok - jumlahBeli

	fmt.Println("--- LAPORAN TRANSAKSI GUDANG ---")
	fmt.Printf("ID Barang      : %d\n", idBarang)
	fmt.Printf("Harga Satuan   : Rp.%.2f\n", hargaSatuan)
	fmt.Printf("Stok Saat ini  : %d\n", stok)
	fmt.Printf("Jumlah Dibeli  : %d\n", jumlahBeli)
	fmt.Println("--------------------------------")
	fmt.Printf("Total Bayar    : Rp.%.2f\n", totalBayar)
	fmt.Printf("Stok Baru      : %d\n", stokBaru)
}

//
// CATATAN TIPE DATA NUMBER & FORMAT PRINT
//
// Tipe data number:
//   - int    : bilangan bulat, ukuran mengikuti platform (32/64 bit). Untuk indeks, loop, umum.
//   - int8   : bilangan bulat 1 byte, rentang -128 sampai 127.
//   - uint   : bilangan bulat non-negatif, ukuran mengikuti platform. Untuk nilai >= 0.
//   - float32: bilangan desimal 4 byte, presisi ~6–7 digit. Untuk hemat memori.
//   - float64: bilangan desimal 8 byte, presisi ~15–16 digit. Default untuk literal desimal di Go.
//
// Format cetak (Printf):
//   - %d   : untuk integer (decimal). Contoh: umur, jumlah, indeks.
//   - \n   : newline (pindah baris). Printf tidak tambah newline otomatis, jadi pakai \n.
//   - %.0f : float tanpa digit di belakang koma (bulat). Cocok untuk kurs rupiah, nominal bulat.
//   - %.2f : float dengan 2 digit di belakang koma. Cocok untuk uang, persen, nilai presisi.
//
