package main

import "fmt"

func main() {
	umur := 20

	// if sederhana
	if umur >= 17 {
		fmt.Println("Sudah dianggap dewasa (>= 17).")
	}

	nilai := 75

	// if - else if - else
	if nilai >= 90 {
		fmt.Println("Nilai: A")
	} else if nilai >= 80 {
		fmt.Println("Nilai: B")
	} else if nilai >= 70 {
		fmt.Println("Nilai: C")
	} else {
		fmt.Println("Nilai: perlu diperbaiki")
	}

	// Inisialisasi di if: variabel hanya hidup di blok if/else if/else ini
	if skor := hitungSkor(8, 9); skor >= 15 {
		fmt.Println("Lulus ujian, skor:", skor)
	} else {
		fmt.Println("Belum lulus, skor:", skor)
	}
	// skor tidak bisa dipakai di sini — scope-nya terbatas pada if di atas

	// Pola umum: cek error setelah operasi
	if err := cobaBagi(10, 2); err != nil {
		fmt.Println("Error:", err)
	}
}

func hitungSkor(a, b int) int {
	return a + b
}

func cobaBagi(a, b int) error {
	if b == 0 {
		return fmt.Errorf("pembagi tidak boleh nol")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
	return nil
}
