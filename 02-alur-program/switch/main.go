package main

import "fmt"

func main() {
	hari := "Sabtu"

	// switch pada nilai
	switch hari {
	case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
		fmt.Println("Hari kerja")
	case "Sabtu", "Minggu":
		fmt.Println("Akhir pekan")
	default:
		fmt.Println("Nama hari tidak dikenal")
	}

	// switch tanpa nilai = switch true (sering untuk rentang atau kondisi panjang)
	umur := 16
	switch {
	case umur < 13:
		fmt.Println("Anak-anak")
	case umur < 20:
		fmt.Println("Remaja")
	default:
		fmt.Println("Dewasa")
	}

	// fallthrough: lanjut ke case berikutnya (jarang dipakai; gunakan hati-hati)
	n := 1
	switch n {
	case 1:
		fmt.Println("satu")
		fallthrough
	case 2:
		fmt.Println("dua (termasuk dari fallthrough)")
	case 3:
		fmt.Println("tiga")
	}
}
