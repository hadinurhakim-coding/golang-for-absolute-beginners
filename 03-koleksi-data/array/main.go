package main

import "fmt"

func main() {
	// Array: ukuran tetap [n]T — bagian dari identitas tipe (bukan seperti slice).
	var angka [3]int
	angka[0] = 10
	angka[1] = 20
	angka[2] = 30
	fmt.Println("array:", angka, "panjang:", len(angka))

	// Literal array
	nama := [3]string{"Andi", "Budi", "Citra"}
	fmt.Println(nama)

	// Compiler hitung panjang dengan [...]
	auto := [...]int{1, 2, 3, 4}
	fmt.Println("auto len:", len(auto))

	// Assign array ke variabel lain = salin seluruh elemen (bukan referensi).
	a := [2]int{1, 2}
	b := a
	b[0] = 99
	fmt.Println("a tetap:", a, "b salinan:", b)

	// Array bisa di-slice menjadi slice yang berbagi backing (lihat materi slice).
	s := a[:] // slice yang melihat data di a
	fmt.Println("slice dari array a:", s)
}
