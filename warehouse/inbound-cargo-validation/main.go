package main

import "fmt"

// Barang merepresentasikan data kiriman inbound (barang masuk) ke gudang.
type Barang struct {
	SKU          string
	JumlahBarang int
	BeratSatuan  float64
	IsFragile    bool
}

func main() {
	// Data sample kiriman kontainer
	barang := Barang{
		SKU:          "WH-ITEM-001",
		JumlahBarang: 50,
		BeratSatuan:  7.5,
		IsFragile:    true,
	}

	// Anonymous function: hitung total berat kiriman (Jumlah × Berat Satuan)
	hitungTotalBerat := func(jumlah int, beratSatuan float64) float64 {
		return float64(jumlah) * beratSatuan
	}
	totalBerat := hitungTotalBerat(barang.JumlahBarang, barang.BeratSatuan)

	// Validasi: SKU tidak kosong
	skuValid := len(barang.SKU) > 0
	// Validasi: jumlah barang > 0
	jumlahValid := barang.JumlahBarang > 0
	// Validasi: berat satuan > 0
	beratValid := barang.BeratSatuan > 0
	// Aturan safety: jika fragile, berat satuan tidak boleh > 10 kg
	fragileSafe := !barang.IsFragile || barang.BeratSatuan <= 10.0

	isApproved := skuValid && jumlahValid && beratValid && fragileSafe

	// Output laporan status
	fmt.Println("--- VALIDASI INBOUND CARGO (BARANG MASUK) ---")
	fmt.Printf("SKU           : %s\n", barang.SKU)
	fmt.Printf("Jumlah Barang : %d\n", barang.JumlahBarang)
	fmt.Printf("Berat Satuan  : %.2f kg\n", barang.BeratSatuan)
	fmt.Printf("Is Fragile    : %t\n", barang.IsFragile)
	fmt.Printf("Total Berat   : %.2f kg\n", totalBerat)
	fmt.Println("---------------------------------------------")
	fmt.Printf("SKU valid (tidak kosong) : %t\n", skuValid)
	fmt.Printf("Jumlah valid (> 0)       : %t\n", jumlahValid)
	fmt.Printf("Berat valid (> 0)       : %t\n", beratValid)
	fmt.Printf("Safety fragile (≤10 kg)  : %t\n", fragileSafe)
	fmt.Println("---------------------------------------------")
	if isApproved {
		fmt.Println("Status : DITERIMA — Barang layak masuk ke sistem.")
	} else {
		fmt.Println("Status : DITOLAK — Barang tidak memenuhi syarat validasi.")
	}
}
