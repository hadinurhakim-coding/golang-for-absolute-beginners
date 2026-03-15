// =============================================================================
// NAMED RETURN VALUE DI GO
// =============================================================================
// Named return value = return value yang diberi nama di signature fungsi.
// Bentuk: func nama() (x int, y string). Variabel x dan y dianggap dideklarasikan
// di awal fungsi (zero value). "Naked return" = return tanpa argumen, mengembalikan
// nilai variabel tersebut. Berguna untuk dokumentasi dan kode singkat di fungsi pendek.
//
// POLA KOMENTAR (konsisten):
//   1. Penjelasan   — ringkasan materi / apa yang dilakukan.
//   2. Alasan penggunaan — kenapa fitur ini dipakai.
//   3. Fungsi penggunaan — cara pakai (pemanggil, assign, naked return).
//   4. Output       — contoh hasil cetakan/return dari fungsi tersebut.
// =============================================================================

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("========== 1. Dasar: named return dan naked return ==========")
	dasarNamedReturn()

	fmt.Println("\n========== 2. Named return diinisialisasi zero value ==========")
	zeroValueNamed()

	fmt.Println("\n========== 3. Bisa tetap return eksplisit ==========")
	returnEksplisit()

	fmt.Println("\n========== 4. Named return untuk dokumentasi ==========")
	dokumentasiNamed()

	fmt.Println("\n========== 5. Studi kasus: min dan max ==========")
	kasusMinMax()

	fmt.Println("\n========== 6. Studi kasus: bagi dan sisa (named) ==========")
	kasusBagiSisa()

	fmt.Println("\n========== 7. Studi kasus: parse nama (depan, belakang) ==========")
	kasusParseNama()

	fmt.Println("\n========== 8. Studi kasus: koordinat (x, y) ==========")
	kasusKoordinat()
}

// -----------------------------------------------------------------------------
// 1. Dasar: named return dan naked return
//
// Penjelasan:
//   Di signature ditulis (hasil int) bukan (int). Variabel "hasil" ada di
//   scope fungsi dan diinisialisasi zero value (0). return tanpa argumen
//   (naked return) mengembalikan nilai current dari "hasil".
//
// Alasan penggunaan:
//   Kode lebih singkat untuk fungsi pendek; nama di signature menjelaskan
//   arti return value tanpa melihat body.
//
// Fungsi penggunaan:
//   x := increment(5). Di dalam: hasil = 6; return (naked). Pemanggil dapat 6.
//
// Output:
//   increment(5) = 6
// -----------------------------------------------------------------------------
func dasarNamedReturn() {
	fmt.Printf("  increment(5) = %d\n", increment(5))
}

func increment(n int) (hasil int) {
	hasil = n + 1
	return
}

// -----------------------------------------------------------------------------
// 2. Named return diinisialisasi zero value
//
// Penjelasan:
//   Variabel named return otomatis diinisialisasi zero value (0 untuk int,
//   "" untuk string, nil untuk pointer/slice/map). Bisa diassign lalu return.
//
// Alasan penggunaan:
//   Untuk return default tanpa menulis return 0, return "", dll. Berguna
//   jika ada beberapa path yang mengubah nilai lalu return di akhir.
//
// Fungsi penggunaan:
//   count := getCount(). Variabel "count" di dalam mulai 0, bisa diubah lalu return.
//
// Output:
//   getCount() = 3
// -----------------------------------------------------------------------------
func zeroValueNamed() {
	fmt.Printf("  getCount() = %d\n", getCount())
}

func getCount() (count int) {
	count = 3
	return
}

// -----------------------------------------------------------------------------
// 3. Bisa tetap return eksplisit
//
// Penjelasan:
//   Meski return value punya nama, tetap boleh menulis return nilai1, nilai2.
//   Nilai yang dikembalikan adalah yang ditulis di return, bukan variabel
//   named (kecuali naked return).
//
// Alasan penggunaan:
//   Kadang satu path pakai naked return, path lain return eksplisit. Atau
//   ingin return literal tanpa assign ke variabel named dulu.
//
// Fungsi penggunaan:
//   a, b := swapNamed(1, 2). Di dalam: return b, a (eksplisit).
//
// Output:
//   swapNamed(1, 2) = 2, 1
// -----------------------------------------------------------------------------
func returnEksplisit() {
	a, b := swapNamed(1, 2)
	fmt.Printf("  swapNamed(1, 2) = %d, %d\n", a, b)
}

func swapNamed(x, y int) (first int, second int) {
	return y, x
}

// -----------------------------------------------------------------------------
// 4. Named return untuk dokumentasi
//
// Penjelasan:
//   Nama di (min int, max int) menjelaskan arti tiap return tanpa perlu
//   komentar tambahan. Pembaca langsung tahu return pertama = min, kedua = max.
//
// Alasan penggunaan:
//   Dokumentasi di signature; mengurangi salah baca urutan return. Terutama
//   berguna untuk multiple return dengan tipe sama.
//
// Fungsi penggunaan:
//   low, high := minMaxNamed(10, 5). Signature (min, max int) menjelaskan urutan.
//
// Output:
//   minMaxNamed(10, 5) = min 5, max 10
// -----------------------------------------------------------------------------
func dokumentasiNamed() {
	mn, mx := minMaxNamed(10, 5)
	fmt.Printf("  minMaxNamed(10, 5) = min %d, max %d\n", mn, mx)
}

func minMaxNamed(a, b int) (min int, max int) {
	if a < b {
		min, max = a, b
	} else {
		min, max = b, a
	}
	return
}

// -----------------------------------------------------------------------------
// 5. Studi kasus: min dan max
//
// Penjelasan:
//   Fungsi minMax(a, b int) (min, max int) memakai named return. Di dalam
//   assign min dan max lalu naked return. Nilai yang dikembalikan jelas dari nama.
//
// Alasan penggunaan:
//   Dua return dengan tipe sama; nama min dan max mencegah tertukar urutan
//   saat pemanggil: low, high := minMax(7, 3).
//
// Fungsi penggunaan:
//   low, high := minMax(7, 3) → low=3, high=7.
//
// Output:
//   minMax(7, 3) = 3, 7
// -----------------------------------------------------------------------------
func kasusMinMax() {
	low, high := minMax(7, 3)
	fmt.Printf("  minMax(7, 3) = %d, %d\n", low, high)
}

func minMax(a, b int) (min int, max int) {
	if a < b {
		min, max = a, b
	} else {
		min, max = b, a
	}
	return
}

// -----------------------------------------------------------------------------
// 6. Studi kasus: bagi dan sisa (named)
//
// Penjelasan:
//   divideNamed(a, b int) (quotient, remainder int). Nama quotient dan
//   remainder menjelaskan arti return. Di dalam: quotient = a/b; remainder = a%b; return.
//
// Alasan penggunaan:
//   Dua nilai integer; tanpa nama pemanggil harus ingat urutan (hasil bagi dulu
//   atau sisa dulu). Named return membuat signature self-documenting.
//
// Fungsi penggunaan:
//   q, r := divideNamed(17, 5) → q=3, r=2.
//
// Output:
//   divideNamed(17, 5) = quotient 3, remainder 2
// -----------------------------------------------------------------------------
func kasusBagiSisa() {
	q, r := divideNamed(17, 5)
	fmt.Printf("  divideNamed(17, 5) = quotient %d, remainder %d\n", q, r)
}

func divideNamed(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

// -----------------------------------------------------------------------------
// 7. Studi kasus: parse nama (depan, belakang)
//
// Penjelasan:
//   splitNamaNamed(s string) (depan, belakang string). Named return untuk
//   dua string. Di dalam assign ke depan dan belakang (dari strings.Fields)
//   lalu naked return.
//
// Alasan penggunaan:
//   Return dua string; nama depan dan belakang menjelaskan urutan dan makna
//   tanpa perlu komentar.
//
// Fungsi penggunaan:
//   d, b := splitNamaNamed("Budi Santoso") → d="Budi", b="Santoso".
//
// Output:
//   splitNamaNamed("Budi Santoso") = depan: Budi, belakang: Santoso
// -----------------------------------------------------------------------------
func kasusParseNama() {
	depan, belakang := splitNamaNamed("Budi Santoso")
	fmt.Printf("  splitNamaNamed(\"Budi Santoso\") = depan: %s, belakang: %s\n", depan, belakang)
}

func splitNamaNamed(s string) (depan string, belakang string) {
	parts := strings.Fields(s)
	if len(parts) >= 2 {
		depan = parts[0]
		belakang = parts[len(parts)-1]
	} else if len(parts) == 1 {
		depan = parts[0]
	}
	return
}

// -----------------------------------------------------------------------------
// 8. Studi kasus: koordinat (x, y)
//
// Penjelasan:
//   getCenter() (x, y int) mengembalikan koordinat pusat. Named return x dan y
//   membuat signature jelas: return pertama = x, kedua = y.
//
// Alasan penggunaan:
//   Koordinat selalu punya konvensi (x, y); nama di signature memastikan
//   pemanggil tidak tertukar dengan tipe lain yang juga (int, int).
//
// Fungsi penggunaan:
//   cx, cy := getCenter() → cx=50, cy=50.
//
// Output:
//   getCenter() = (50, 50)
// -----------------------------------------------------------------------------
func kasusKoordinat() {
	x, y := getCenter()
	fmt.Printf("  getCenter() = (%d, %d)\n", x, y)
}

func getCenter() (x int, y int) {
	x = 50
	y = 50
	return
}
