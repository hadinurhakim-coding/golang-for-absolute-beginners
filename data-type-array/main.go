package main

import "fmt"

func main() {
	// 1. Array: deklarasi [n]T — ukuran tetap, bagian dari tipe
	var stokPerRak [5]int
	stokPerRak[0] = 10
	stokPerRak[1] = 25
	stokPerRak[2] = 8
	stokPerRak[3] = 0
	stokPerRak[4] = 42

	fmt.Println("--- ARRAY STOK PER RAK (5 rak) ---")
	fmt.Printf("Stok per rak : %v\n", stokPerRak)
	fmt.Printf("Jumlah rak   : %d (len)\n", len(stokPerRak))
	fmt.Printf("Rak ke-3     : %d\n", stokPerRak[2])
	fmt.Println("----------------------------------")

	fmt.Println()

	// 2. Array literal — inisialisasi sekaligus
	kodeBarang := [4]string{"BRG-001", "BRG-002", "BRG-003", "BRG-004"}
	fmt.Println("--- ARRAY KODE BARANG ---")
	for i := 0; i < len(kodeBarang); i++ {
		fmt.Printf("  [%d] %s\n", i, kodeBarang[i])
	}
	fmt.Println("-------------------------")

	fmt.Println()

	// 3. [...] — compiler tentukan ukuran dari jumlah elemen
	jumlahKeluarPerHari := [...]int{50, 30, 45, 20, 60, 10, 0}
	fmt.Println("--- KELUAR BARANG PER HARI (Senin–Minggu) ---")
	fmt.Printf("Data : %v\n", jumlahKeluarPerHari)
	fmt.Printf("Len  : %d\n", len(jumlahKeluarPerHari))
	fmt.Println("---------------------------------------------")

	fmt.Println()

	// 4. Zero value — array belum diisi bernilai nol untuk tipe elemen
	var kapasitasRak [3]int
	fmt.Println("--- ZERO VALUE ARRAY ---")
	fmt.Printf("kapasitasRak (belum diisi) : %v\n", kapasitasRak)
	fmt.Println("-------------------------")

	fmt.Println()

	// 5. Iterasi dengan range
	totalStok := 0
	for _, jumlah := range stokPerRak {
		totalStok += jumlah
	}
	fmt.Println("--- TOTAL STOK (range) ---")
	fmt.Printf("Total stok semua rak : %d\n", totalStok)
	fmt.Println("--------------------------")
}

//
// CATATAN TIPE DATA ARRAY DI GO
//
// Bentuk: [n]T — n = panjang tetap (tipe), T = tipe elemen. Ukuran bagian dari tipe; [3]int dan [5]int berbeda tipe.
//
// Deklarasi: var a [n]T (zero value), b := [n]T{v1, v2, ...}, c := [...]T{v1, v2, ...} (ukuran dari jumlah elemen).
//
// Operasi: a[i] (baca/tulis), len(a). Tidak ada cap (array selalu fixed). Iterasi: for i := range a, for i, v := range a.
//
// Sifat: array adalah value type — assign atau pass ke fungsi = copy, bukan referensi. Untuk koleksi dinamis pakai slice ([]T).
//
// Lihat juga: slice-array (slice vs array), data-type-number, data-type-string.
//
