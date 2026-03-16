package main

import "fmt"

type IDBarang uint32
type KodeGudang string
type JumlahStok int
type Rupiah = float64

func main() {
	var id IDBarang = 1001
	kode := KodeGudang("KIM-JKT-01")
	stok := JumlahStok(100)

	fmt.Println("--- TYPE DECLARATION — SISTEM GUDANG ---")
	fmt.Printf("ID Barang : %d\n", id)
	fmt.Printf("Kode Gudang : %s\n", kode)
	fmt.Printf("Stok : %d\n", stok)
	fmt.Println("----------------------------------------")

	fmt.Println()

	var hargaSatuan Rupiah = 20000.25
	totalNilai := hargaSatuan * Rupiah(stok)
	fmt.Printf("harga Satuan : Rp.%.2f\n", hargaSatuan)
	fmt.Printf("Total Nilai : Rp.%.2f\n", totalNilai)

	fmt.Println()

	var isAsUint uint32 = uint32(id)
	fmt.Printf("ID Barang sebagai uint32 : %d\n", isAsUint)
}
