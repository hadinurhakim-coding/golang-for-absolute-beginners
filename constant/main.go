package main

import "fmt"

func main() {
	// 1. Konstanta tunggal — nilai tetap untuk seluruh program
	const namaPerusahaan = "PT. Gudang KIM"
	const pajakPersen = 11
	const kapasitasMaxGudang = 10000

	fmt.Println("--- KONSTANTA SISTEM GUDANG ---")
	fmt.Printf("Nama Perusahaan : %s\n", namaPerusahaan)
	fmt.Printf("Tarif Pajak    : %d%%\n", pajakPersen)
	fmt.Printf("Kapasitas Max  : %d\n", kapasitasMaxGudang)
	fmt.Println("--------------------------------")

	fmt.Println()

	// 2. Blok const — mengelompokkan konstanta terkait (seperti var block)
	const (
		statusBarangKosong   = 0
		statusBarangTersedia = 1
		statusBarangRusak    = 2
	)

	statusStok := statusBarangTersedia
	fmt.Printf("Status stok saat ini : %d\n", statusStok)
	if statusStok == statusBarangTersedia {
		fmt.Println("Barang siap dikirim.")
	}

	fmt.Println()

	// 3. Konstanta bertipe vs tidak bertipe
	// Tanpa tipe (untyped): bisa dipakai dalam ekspresi dengan tipe berbeda selama kompatibel
	const diskon = 10
	var harga float64 = 100000
	potongan := harga * diskon / 100
	fmt.Printf("Harga awal : Rp.%.0f\n", harga)
	fmt.Printf("Diskon %d%% : Rp.%.0f\n", diskon, potongan)
	fmt.Printf("Harga akhir : Rp.%.0f\n", harga-potongan)

	fmt.Println()

	// 4. Konstanta bertipe — tipe eksplisit, hanya dipakai di konteks tipe yang sama
	const kodeGudang string = "WH-JKT-01"
	const batasStokMin int = 5
	fmt.Printf("Kode Gudang    : %s\n", kodeGudang)
	fmt.Printf("Batas Stok Min : %d\n", batasStokMin)
}

//
// CATATAN KONSTANTA DI GO
//
// Deklarasi:
//   - const nama = nilai        : konstanta untyped; tipe disimpulkan dari konteks (number, string, bool).
//   - const nama tipe = nilai   : konstanta bertipe; harus dipakai sesuai tipe tersebut.
//   - const ( ... )             : blok const untuk beberapa konstanta terkait (status, kode, limit).
//
// Perbedaan dengan variabel:
//   - Konstanta tidak bisa di-reassign (nilainya tetap). Untuk nilai yang bisa berubah pakai var atau :=.
//   - Konstanta bisa dideklarasikan di tingkat package (global) atau di dalam fungsi.
//
// Format cetak (Printf):
//   - %d   : integer. Untuk konstanta numerik (status, persen, kapasitas).
//   - %s   : string. Untuk teks (nama perusahaan, kode gudang).
//   - %.0f / %.2f : float. Untuk uang (harga, potongan). \n untuk newline.
//   - %%   : literal karakter %. Contoh: "Tarif Pajak : %d%%" → menampilkan angka lalu simbol %.
//
// Lihat juga: variable (var vs const), data-type-number, data-type-string, data-type-boolean.
//
