package main

import "fmt"

func main() {
	// 1. Variabel dengan tipe data dan nilai awal (tidak wajib)
	var namaGudang string = "Gudang Milik KIM"
	var kapasitasGudang int
	kapasitasGudang = 10000

	// 2. idiomatis untuk deklarasi variabel tanpa tipe data
	// menggunakan operator := untuk deklarasi variabel tanpa tipe data
	lokasi := "Jl. KIM Factory, indonesia"
	jumlahBarang := 20000
	isAktif := true

	// Output informasi gudang
	fmt.Println("---INFORMASI GUDANG---")
	fmt.Println("Nama Gudang :", namaGudang)
	fmt.Println("Lokasi Gudang :", lokasi)
	fmt.Println("Kapasitas Gudang :", kapasitasGudang)
	fmt.Println("Jumlah Barang :", jumlahBarang)
	fmt.Println("Status Aktif :", isAktif)

	//Baris
	fmt.Println()

	//3. Blok variabel untuk deklarasi multi variabel
	// stok barang
	var (
		idBarang     uint32  = 1001
		stokAwal     int     = 50
		hargaSatuan  float64 = 15000.64
		barangKeluar int     = 10
	)

	// hitung total nilai barang
	totalNilai := float64(stokAwal) * hargaSatuan

	// Output laporan stok barang
	fmt.Println("---LAPORAN STOK GUDANG---")
	fmt.Printf("ID BARANG    : %d\n", idBarang)
	fmt.Printf("STOK AWAL    : %d\n", stokAwal)
	fmt.Printf("BARANG KELUAR: %d\n", barangKeluar)
	fmt.Printf("HARGA SATUAN : Rp.%.2f\n", hargaSatuan)
	fmt.Printf("TOTAL NILAI  : Rp.%.2f\n", totalNilai)
	fmt.Println("--------------------------------")

	fmt.Println()

	stokAkhir := stokAwal - barangKeluar
	fmt.Printf("STOK AKHIR  : %d\n", stokAkhir)
}

//
// CATATAN VARIABEL DI GO
//
// Cara deklarasi:
//   - var nama tipe [= nilai] : deklarasi eksplisit. Jika nilai tidak diberikan, dipakai zero value ("" untuk string, 0 untuk number, false untuk bool).
//   - nama := nilai           : short declaration di dalam fungsi; tipe disimpulkan dari nilai. Idiomatik untuk inisialisasi.
//
// Pengelompokan:
//   - var ( ... ) : blok var untuk beberapa variabel terkait; berguna agar kode rapi dan alignment jelas.
//
// Sifat:
//   - Variabel bisa di-reassign (nilainya diubah). Untuk nilai tetap pakai const (lihat constant/).
//   - Naming: camelCase untuk variabel (namaGudang, stokAwal). Export pakai PascalCase (NamaGudang).
//
// Format cetak (Printf — dipakai di file ini):
//   - Println("teks", x) : cetak teks dan nilai x, dipisah spasi, lalu newline. Tidak ada format verb; nilai dicetak apa adanya.
//   - Printf("template", args...) : template bisa berisi format verb; setiap verb diganti oleh argumen yang sesuai.
//   - %d   : integer (decimal). Untuk int, int32, uint32, dll. Contoh: ID barang, stok, jumlah.
//   - %.2f : float dengan 2 digit di belakang koma. Cocok untuk uang (harga, total nilai). Angka dibulatkan ke 2 desimal.
//   - \n   : newline (pindah baris). Printf tidak menambah newline otomatis, jadi pakai \n di akhir jika perlu.
//
// Lihat juga: data-type-number, data-type-string, data-type-boolean, format-verb, constant.
//
