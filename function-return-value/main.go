// =============================================================================
// FUNCTION RETURN VALUE DI GO
// =============================================================================
// Return value = nilai yang dikembalikan fungsi ke pemanggil. Bentuk: func nama() T
// atau func nama() (T1, T2). Pakai return v atau return v1, v2. Pemanggil terima
// dengan assign: x := f() atau a, b := f(). Bisa ignore dengan _.
//
// POLA KOMENTAR (konsisten):
//   1. Penjelasan   — ringkasan materi / apa yang dilakukan.
//   2. Alasan penggunaan — kenapa fitur ini dipakai.
//   3. Fungsi penggunaan — cara pakai (pemanggil, assign, ignore).
//   4. Output       — contoh hasil cetakan/return dari fungsi tersebut.
// =============================================================================

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("========== 1. Satu return value ==========")
	satuReturn()

	fmt.Println("\n========== 2. Multiple return values ==========")
	multipleReturn()

	fmt.Println("\n========== 3. Named return values ==========")
	namedReturn()

	fmt.Println("\n========== 4. Ignore return value dengan _ ==========")
	ignoreReturn()

	fmt.Println("\n========== 5. Return early (return sebelum akhir fungsi) ==========")
	returnEarly()

	fmt.Println("\n========== 6. Return pointer (*T) ==========")
	returnPointer()

	fmt.Println("\n========== 7. Studi kasus: bagi dan sisa (multiple return) ==========")
	kasusBagiSisa()

	fmt.Println("\n========== 8. Studi kasus: parse string ke int (value + error) ==========")
	kasusParseInt()

	fmt.Println("\n========== 9. Studi kasus: min dan max (named return) ==========")
	kasusMinMax()
}

// -----------------------------------------------------------------------------
// 1. Satu return value
//
// Penjelasan:
//   Fungsi mengembalikan tepat satu nilai. Tipe return ditulis setelah ( ):
//   func nama() T. Di dalam fungsi pakai return nilai. Pemanggil: x := nama().
//
// Alasan penggunaan:
//   Hasil satu nilai (hasil hitung, status, string hasil) dipakai langsung
//   atau disimpan ke variabel.
//
// Fungsi penggunaan:
//   hasil := double(5). Return value ditangkap ke variabel atau dipakai
//   di ekspresi: fmt.Println(double(5)).
//
// Output:
//   double(5) = 10
//   greet("Budi") = Halo, Budi
// -----------------------------------------------------------------------------
func satuReturn() {
	fmt.Printf("  double(5) = %d\n", double(5))
	fmt.Printf("  greet(\"Budi\") = %s\n", greetReturn("Budi"))
}

func double(x int) int {
	return x * 2
}

func greetReturn(nama string) string {
	return "Halo, " + nama
}

// -----------------------------------------------------------------------------
// 2. Multiple return values
//
// Penjelasan:
//   Fungsi bisa mengembalikan lebih dari satu nilai: func nama() (T1, T2).
//   Return: return v1, v2. Pemanggil: a, b := nama(). Urutan harus sesuai.
//
// Alasan penggunaan:
//   Pola umum di Go: hasil + error (value, err). Juga pasangan nilai seperti
//   hasil bagi dan sisa, koordinat x,y, min dan max, tanpa perlu struct.
//
// Fungsi penggunaan:
//   q, r := divide(10, 3). Bisa ignore salah satu: v, _ := divide(10, 3).
//
// Output:
//   divide(10, 3) = 3 sisa 1
// -----------------------------------------------------------------------------
func multipleReturn() {
	q, r := divide(10, 3)
	fmt.Printf("  divide(10, 3) = %d sisa %d\n", q, r)
}

func divide(a, b int) (int, int) {
	return a / b, a % b
}

// -----------------------------------------------------------------------------
// 3. Named return values
//
// Penjelasan:
//   Return value bisa diberi nama di signature: func nama() (x int, y int).
//   Variabel x, y dianggap dideklarasikan di awal fungsi. return tanpa
//   argumen (naked return) mengembalikan nilai variabel tersebut.
//
// Alasan penggunaan:
//   Dokumentasi di signature; kode kadang lebih singkat. Berguna untuk return
//   yang banyak atau nilai default yang jelas.
//
// Fungsi penggunaan:
//   min, max := minMax(7, 3). Di dalam minMax bisa hanya return (naked).
//
// Output:
//   minMax(7, 3) = min 3, max 7
// -----------------------------------------------------------------------------
func namedReturn() {
	mn, mx := minMax(7, 3)
	fmt.Printf("  minMax(7, 3) = min %d, max %d\n", mn, mx)
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
// 4. Ignore return value dengan _
//
// Penjelasan:
//   Pemanggil bisa mengabaikan salah satu (atau semua) return value dengan
//   blank identifier _. Contoh: v, _ := f() hanya mengambil nilai pertama.
//
// Alasan penggunaan:
//   Saat salah satu return tidak dibutuhkan (misalnya error yang sengaja
//   diabaikan, atau indeks dari range). Menghindari variabel tidak terpakai.
//
// Fungsi penggunaan:
//   q, _ := divide(10, 3) hanya menyimpan hasil bagi. _, r := divide(10, 3)
//   hanya menyimpan sisa.
//
// Output:
//   Hanya quotient: 3
//   Hanya remainder: 1
// -----------------------------------------------------------------------------
func ignoreReturn() {
	q, _ := divide(10, 3)
	fmt.Printf("  Hanya quotient: %d\n", q)
	_, r := divide(10, 3)
	fmt.Printf("  Hanya remainder: %d\n", r)
}

// -----------------------------------------------------------------------------
// 5. Return early
//
// Penjelasan:
//   return bisa dipanggil di mana saja dalam fungsi. Begitu return dieksekusi,
//   fungsi selesai dan nilai dikembalikan. Tidak wajib return di akhir saja.
//
// Alasan penggunaan:
//   Menghindari nested if; keluar segera saat kondisi terpenuhi (guard clause)
//   atau nilai default sudah bisa dikembalikan.
//
// Fungsi penggunaan:
//   Di dalam fungsi: if n < 0 { return 0 }; ... return n*2. Pemanggil tidak
//   berubah: x := abs(-5).
//
// Output:
//   abs(-5) = 5
//   abs(3) = 3
// -----------------------------------------------------------------------------
func returnEarly() {
	fmt.Printf("  abs(-5) = %d\n", abs(-5))
	fmt.Printf("  abs(3) = %d\n", abs(3))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// -----------------------------------------------------------------------------
// 6. Return pointer (*T)
//
// Penjelasan:
//   Fungsi bisa mengembalikan alamat (pointer): func nama() *T. Return &v.
//   Pemanggil dapat *T dan bisa mengubah nilai yang ditunjuk. Aman mengembalikan
//   pointer ke variabel lokal karena Go menganalisis escape ke heap.
//
// Alasan penggunaan:
//   Mengembalikan struct besar tanpa copy; atau agar pemanggil bisa memodifikasi
//   nilai (shared state). Juga untuk "optional" (nil = tidak ada).
//
// Fungsi penggunaan:
//   p := newInt(42); *p == 42. Pemanggil pakai *p untuk baca/ubah.
//
// Output:
//   newInt(42) = 42
// -----------------------------------------------------------------------------
func returnPointer() {
	p := newInt(42)
	fmt.Printf("  newInt(42) = %d\n", *p)
}

func newInt(x int) *int {
	v := x
	return &v
}

// -----------------------------------------------------------------------------
// 7. Studi kasus: bagi dan sisa (multiple return)
//
// Penjelasan:
//   Fungsi divide(a, b int) (quotient, remainder int) mengembalikan hasil
//   bagi bulat dan sisa bagi. Menggambarkan multiple return untuk dua nilai
//   terkait.
//
// Alasan penggunaan:
//   Satu operasi (pembagian) menghasilkan dua informasi; multiple return
//   menghindari struct atau parameter pointer hanya untuk "keluaran" kedua.
//
// Fungsi penggunaan:
//   q, r := divide(17, 5) → q=3, r=2.
//
// Output:
//   divide(17, 5) = 3 sisa 2
// -----------------------------------------------------------------------------
func kasusBagiSisa() {
	q, r := divide(17, 5)
	fmt.Printf("  divide(17, 5) = %d sisa %d\n", q, r)
}

// -----------------------------------------------------------------------------
// 8. Studi kasus: parse string ke int (value + error)
//
// Penjelasan:
//   Pola umum di Go: fungsi mengembalikan (value T, err error). Jika sukses
//   err == nil; jika gagal value biasanya zero value dan err berisi penyebab.
//   Di sini pakai strconv.Atoi yang return (int, error).
//
// Alasan penggunaan:
//   Error handling eksplisit tanpa exception. Pemanggil wajib cek err dan
//   memutuskan tindakan (retry, log, return error ke atas).
//
// Fungsi penggunaan:
//   n, err := strconv.Atoi("42"); if err != nil { ... }; else pakai n.
//
// Output:
//   Parse "42" = 42 (sukses)
//   Parse "abc" = 0 (error: ...)
// -----------------------------------------------------------------------------
func kasusParseInt() {
	n, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("  Parse \"42\" error: %v\n", err)
	} else {
		fmt.Printf("  Parse \"42\" = %d (sukses)\n", n)
	}
	n2, err2 := strconv.Atoi("abc")
	if err2 != nil {
		fmt.Printf("  Parse \"abc\" = %d (error: %v)\n", n2, err2)
	} else {
		fmt.Printf("  Parse \"abc\" = %d (sukses)\n", n2)
	}
}

// -----------------------------------------------------------------------------
// 9. Studi kasus: min dan max (named return)
//
// Penjelasan:
//   Fungsi minMax(a, b int) (min, max int) menggunakan named return. Di dalam
//   fungsi hanya assign ke min/max lalu return. Nilai yang dikembalikan jelas
//   dari nama variabel.
//
// Alasan penggunaan:
//   Signature menjelaskan arti tiap return; naked return membuat kode singkat
//   untuk fungsi pendek.
//
// Fungsi penggunaan:
//   low, high := minMax(10, 5) → low=5, high=10.
//
// Output:
//   minMax(10, 5) = min 5, max 10
// -----------------------------------------------------------------------------
func kasusMinMax() {
	low, high := minMax(10, 5)
	fmt.Printf("  minMax(10, 5) = min %d, max %d\n", low, high)
}
