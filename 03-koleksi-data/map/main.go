package main

import "fmt"

func main() {
	// Map: pasangan kunci-nilai; kunci harus bisa dibandingkan (==, comparable).

	// Literal
	skor := map[string]int{
		"Andi": 80,
		"Budi": 90,
	}
	fmt.Println("literal:", skor)

	// make — berguna saat isi ditambah bertahap
	umur := make(map[string]int)
	umur["Citra"] = 25
	umur["Dedi"] = 30
	fmt.Println("make:", umur)

	// Baca kunci yang ada
	fmt.Println("Andi:", skor["Andi"])

	// Kunci tidak ada: nilai zero tipe nilai (int → 0) — membingungkan tanpa cek.
	fmt.Println("tidak ada:", skor["Eka"]) // 0

	// Comma-ok: bedakan "tidak ada" vs "nilai memang 0"
	if v, ok := skor["Eka"]; ok {
		fmt.Println("ketemu:", v)
	} else {
		fmt.Println("kunci Eka tidak ada")
	}

	// Hapus
	delete(skor, "Andi")
	fmt.Println("setelah delete Andi:", skor)

	// Iterasi (urutan tidak dijamin tetap antar run)
	fmt.Println("--- range map ---")
	for nama, nilai := range skor {
		fmt.Println(nama, nilai)
	}
}
