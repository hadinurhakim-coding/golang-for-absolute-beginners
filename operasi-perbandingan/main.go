// =============================================================================
// OPERASI PERBANDINGAN DI GO
// =============================================================================
// Operator: == (sama), != (tidak sama), < (kurang dari), > (lebih dari), <=, >=
// Hasil selalu bool (true/false). Dipakai di if, for, dan ekspresi.
// =============================================================================

package main

import "fmt"

func main() {
	fmt.Println("========== 1. Dasar: angka ==========")
	dasarAngka()

	fmt.Println("\n========== 2. String: ==, !=, <, > ==========")
	perbandinganString()

	fmt.Println("\n========== 3. Float: hati-hati presisi ==========")
	perbandinganFloat()

	fmt.Println("\n========== 4. Contoh kasus: nilai lulus ==========")
	kasusNilaiLulus()

	fmt.Println("\n========== 5. Contoh kasus: range usia ==========")
	kasusRangeUsia()

	fmt.Println("\n========== 6. Contoh kasus: validasi input ==========")
	kasusValidasi()
}

func dasarAngka() {
	a, b := 10, 20
	fmt.Printf("a=%d, b=%d\n", a, b)
	fmt.Printf("a == b → %t\n", a == b)
	fmt.Printf("a != b → %t\n", a != b)
	fmt.Printf("a < b  → %t\n", a < b)
	fmt.Printf("a > b  → %t\n", a > b)
	fmt.Printf("a <= 10 → %t\n", a <= 10)
	fmt.Printf("a >= 10 → %t\n", a >= 10)
}

func perbandinganString() {
	s1, s2 := "apple", "banana"
	fmt.Printf("%q == %q → %t\n", s1, s2, s1 == s2)
	fmt.Printf("%q != %q → %t\n", s1, s2, s1 != s2)
	fmt.Printf("%q < %q → %t (urutan leksikografis/alfabet)\n", s1, s2, s1 < s2)
	fmt.Printf("%q > %q → %t\n", s1, s2, s1 > s2)
}

func perbandinganFloat() {
	f := 0.1 + 0.2
	fmt.Printf("0.1 + 0.2 = %v\n", f)
	fmt.Printf("f == 0.3 → %t (bisa false karena presisi)\n", f == 0.3)
	// Untuk perbandingan float, sering pakai selang (epsilon) atau math.Abs(a-b) < epsilon
}

func kasusNilaiLulus() {
	nilai := 75
	batasLulus := 60
	lulus := nilai >= batasLulus
	fmt.Printf("Nilai: %d, Batas lulus: %d\n", nilai, batasLulus)
	fmt.Printf("Lulus? %t\n", lulus)

	nilai2 := 55
	fmt.Printf("Nilai %d lulus? %t\n", nilai2, nilai2 >= batasLulus)
}

func kasusRangeUsia() {
	usia := 25
	minDewasa, maxRemaja := 18, 17
	dewasa := usia >= minDewasa
	remaja := usia >= 0 && usia <= maxRemaja
	fmt.Printf("Usia %d: dewasa? %t, remaja? %t\n", usia, dewasa, remaja)
}

func kasusValidasi() {
	username := "admin"
	panjangMin := 5
	validPanjang := len(username) >= panjangMin
	fmt.Printf("Username %q, panjang minimal %d: valid? %t\n", username, panjangMin, validPanjang)

	angka := 100
	dalamRange := angka >= 1 && angka <= 100
	fmt.Printf("Angka %d dalam range [1,100]? %t\n", angka, dalamRange)
}
