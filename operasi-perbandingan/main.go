package main

import "fmt"

func main() {
	stokSaatIni := 1001
	kapasitasMax := 10000
	suhuRuangan := 25.5
	suhuIdeal := 20.0
	kodeAkses := "GUDANG-01"
	inputAkses := "GUDANG-B1"

	isAksesValid := kodeAkses == inputAkses
	isStokPenuh := stokSaatIni == kapasitasMax
	isBukanSuhuIdeal := suhuRuangan >= suhuIdeal

	overheat := suhuRuangan > 30.0
	stokKritis := stokSaatIni < 10
	keamanan := isAksesValid && isStokPenuh && !overheat && !stokKritis

	bisaTambahStok := stokSaatIni <= 80
	kapasitasAman := stokSaatIni >= 0

	fmt.Println("=== SISTEM MONITORING GUDANG ===")
	fmt.Printf("Kode Akses Valid: %t\n", isAksesValid)
	fmt.Printf("Status Stok : %t\n", isStokPenuh)
	fmt.Printf("Status Suhu : %t\n", isBukanSuhuIdeal)
	fmt.Printf("Status Overheat : %t\n", overheat)
	fmt.Printf("Status Stok Kritis : %t\n", stokKritis)
	fmt.Printf("Status Keamanan : %t\n", keamanan)
	fmt.Printf("Bisa Tambah Stok : %t\n", bisaTambahStok)
	fmt.Printf("Kapasitas Aman : %t\n", kapasitasAman)
	fmt.Println("===================================")

	if keamanan {
		fmt.Println("Gudang Aman, tetap bekerja.")
	} else {
		fmt.Println("Gudang tidak aman, perlu perbaikan.")
	}

	if stokSaatIni < kapasitasMax {
		fmt.Printf("Masih bisa ditambahkan %d barang lagi.\n", kapasitasMax-stokSaatIni)
	} else {
		fmt.Printf("Kapasitas max %d, stok saat ini %d sudah penuh, perlu dikeluarkan!\n", kapasitasMax, stokSaatIni)
	}

}

//
// CATATAN OPERASI PERBANDINGAN DI GO
//
// Operator: == (sama), != (tidak sama), < (kurang dari), <= (kurang/sama),
//           > (lebih dari), >= (lebih/sama). Hasil selalu bool (true/false).
//
// Tipe: Bandingkan tipe yang sama. int dengan float64 harus di-cast dulu, misal: float64(stok) > suhu.
// String: case sensitive; "Gudang" == "gudang" → false.
//
// Format cetak: %t (bool), %d (int). \n untuk newline di akhir Printf.
//
// Lihat juga: operasi-matematika, operasi-boolean, data-type-boolean, if-statement.
//
