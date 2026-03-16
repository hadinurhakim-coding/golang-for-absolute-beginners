package main

import "fmt"

func main() {
	// 1. Operasi dasar: +, -, *, / — perhitungan stok dan nilai
	stokAwal := 100
	barangMasuk := 25
	barangKeluar := 30
	stokAkhir := stokAwal + barangMasuk - barangKeluar

	hargaSatuan := 15000.0
	jumlahBeli := 5
	totalBayar := hargaSatuan * float64(jumlahBeli)

	fmt.Println("--- OPERASI DASAR (+, -, *, /) ---")
	fmt.Printf("Stok Awal    : %d\n", stokAwal)
	fmt.Printf("Barang Masuk : +%d\n", barangMasuk)
	fmt.Printf("Barang Keluar: -%d\n", barangKeluar)
	fmt.Printf("Stok Akhir   : %d\n", stokAkhir)
	fmt.Println("--------------------------------")
	fmt.Printf("Harga Satuan : Rp.%.0f x %d = Rp.%.0f\n", hargaSatuan, jumlahBeli, totalBayar)
	fmt.Println()

	// 2. Pembagian integer vs float — bagi lusin dan rata-rata
	totalItem := 47
	isiPerKarton := 12
	jumlahKarton := totalItem / isiPerKarton
	sisaItem := totalItem % isiPerKarton

	totalNilai := 450000.0
	jumlahTransaksi := 3
	rataRataNilai := totalNilai / float64(jumlahTransaksi)

	fmt.Println("--- PEMBAGIAN & MODULO ---")
	fmt.Printf("Total Item     : %d\n", totalItem)
	fmt.Printf("Isi per Karton : %d\n", isiPerKarton)
	fmt.Printf("Jumlah Karton   : %d (pembagian integer)\n", jumlahKarton)
	fmt.Printf("Sisa Item      : %d (modulo %%)\n", sisaItem)
	fmt.Println("--------------------------------")
	fmt.Printf("Total Nilai / %d transaksi = Rp.%.0f (rata-rata)\n", jumlahTransaksi, rataRataNilai)
	fmt.Println()

	// 3. Augmented assignment (+=, -=, *=, /=) — update stok dan diskon
	stok := 50
	stok += 20
	stok -= 5

	harga := 100000.0
	diskonPersen := 10.0
	harga *= (1 - diskonPersen/100)

	fmt.Println("--- AUGMENTED ASSIGNMENT (+=, -=, *=) ---")
	fmt.Printf("Stok awal 50, +20 lalu -5 = %d\n", stok)
	fmt.Printf("Harga Rp.100.000 diskon 10%% = Rp.%.0f\n", harga)
	fmt.Println()

	// 4. Ekspresi gabungan — perhitungan satu baris
	quantity := 10
	unitPrice := 25000.0
	pajak := 0.11
	total := float64(quantity) * unitPrice * (1 + pajak)

	fmt.Println("--- EKSPRESI GABUNGAN ---")
	fmt.Printf("Qty %d x Rp.%.0f x (1 + Pajak 11%%) = Rp.%.0f\n", quantity, unitPrice, total)
}

//
// CATATAN OPERASI MATEMATIKA DI GO
//
// Operator aritmetika:
//   - +  : penjumlahan. - : pengurangan. * : perkalian. / : pembagian. %% : sisa bagi (modulo, hanya integer).
//
// Pembagian:
//   - int / int   : hasil integer (truncate). Contoh: 47/12 = 3.
//   - float / float (atau salah satu float): hasil float. Cast: float64(angka) agar pembagian desimal.
//
// Augmented assignment:
//   - +=, -=, *=, /=, %%= : ubah variabel in-place. Contoh: stok += 20 sama dengan stok = stok + 20.
//
// Format cetak: %%d (integer), %%.0f / %%.2f (float). %% untuk literal %%. \\n newline.
//
// Lihat juga: data-type-number, variable, constant, operasi-perbandingan, operasi-boolean.
//
