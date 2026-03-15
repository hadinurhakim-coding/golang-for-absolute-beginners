// =============================================================================
// RETURNING MULTIPLE VALUE DI GO
// =============================================================================
// Go mengizinkan fungsi mengembalikan lebih dari satu nilai. Bentuk:
// func nama() (T1, T2) atau (T1, T2, T3). Return: return v1, v2. Pemanggil
// tangkap dengan assign: a, b := nama(). Pola umum: (result, error).
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
	"strings"
)

func main() {
	fmt.Println("========== 1. Dasar: dua return value ==========")
	duaReturnValue()

	fmt.Println("\n========== 2. Tiga atau lebih return value ==========")
	tigaReturnValue()

	fmt.Println("\n========== 3. Pola value + error (idiomatik Go) ==========")
	polaValueError()

	fmt.Println("\n========== 4. Ignore sebagian return dengan _ ==========")
	ignoreSebagianReturn()

	fmt.Println("\n========== 5. Named return untuk multiple value ==========")
	namedMultipleReturn()

	fmt.Println("\n========== 6. Studi kasus: bagi dan sisa ==========")
	kasusBagiSisa()

	fmt.Println("\n========== 7. Studi kasus: parse nama (depan, belakang) ==========")
	kasusParseNama()

	fmt.Println("\n========== 8. Studi kasus: statistik slice (sum, avg) ==========")
	kasusStatistikSlice()

	fmt.Println("\n========== 9. Studi kasus: lookup map (value, ok) ==========")
	kasusLookupMap()
}

// -----------------------------------------------------------------------------
// 1. Dasar: dua return value
//
// Penjelasan:
//   Fungsi mengembalikan dua nilai. Signature: func nama() (T1, T2). Di dalam
//   fungsi: return v1, v2. Pemanggil: a, b := nama(). Urutan harus sama.
//
// Alasan penggunaan:
//   Satu pemanggilan menghasilkan dua informasi (hasil bagi + sisa, koordinat,
//   min+max) tanpa struct atau parameter pointer untuk "output" kedua.
//
// Fungsi penggunaan:
//   x, y := getPoint(). Nilai pertama ke x, kedua ke y.
//
// Output:
//   getPoint() = (10, 20)
// -----------------------------------------------------------------------------
func duaReturnValue() {
	x, y := getPoint()
	fmt.Printf("  getPoint() = (%d, %d)\n", x, y)
}

func getPoint() (int, int) {
	return 10, 20
}

// -----------------------------------------------------------------------------
// 2. Tiga atau lebih return value
//
// Penjelasan:
//   Go mengizinkan lebih dari dua return: func nama() (T1, T2, T3). Return
//   dan assign dengan jumlah yang sama. Jarang lebih dari 2–3 agar tetap jelas.
//
// Alasan penggunaan:
//   Saat satu operasi menghasilkan beberapa nilai terkait (misalnya sum, min,
//   max) tanpa membuat struct atau banyak pemanggilan terpisah.
//
// Fungsi penggunaan:
//   a, b, c := getTiga(). Semua nilai harus ditangkap atau di-ignore dengan _.
//
// Output:
//   getTiga() = 1, 2, 3
// -----------------------------------------------------------------------------
func tigaReturnValue() {
	a, b, c := getTiga()
	fmt.Printf("  getTiga() = %d, %d, %d\n", a, b, c)
}

func getTiga() (int, int, int) {
	return 1, 2, 3
}

// -----------------------------------------------------------------------------
// 3. Pola value + error (idiomatik Go)
//
// Penjelasan:
//   Konvensi di Go: fungsi yang bisa gagal mengembalikan (value T, err error).
//   Jika sukses, err == nil. Jika gagal, value biasanya zero value dan err
//   berisi penyebab. Pemanggil wajib cek err sebelum pakai value.
//
// Alasan penggunaan:
//   Error handling eksplisit tanpa exception. Alur sukses/gagal jelas di
//   kode; tidak ada panic tersembunyi. Standar di perpustakaan Go.
//
// Fungsi penggunaan:
//   v, err := parseSomething(s); if err != nil { ... }; pakai v.
//
// Output:
//   parseNumber("42") = 42 <nil>
//   parseNumber("x") = 0 error: invalid
// -----------------------------------------------------------------------------
func polaValueError() {
	v1, err1 := parseNumber("42")
	fmt.Printf("  parseNumber(\"42\") = %d %v\n", v1, err1)
	v2, err2 := parseNumber("x")
	fmt.Printf("  parseNumber(\"x\") = %d %v\n", v2, err2)
}

func parseNumber(s string) (int, error) {
	if s == "42" {
		return 42, nil
	}
	return 0, fmt.Errorf("invalid")
}

// -----------------------------------------------------------------------------
// 4. Ignore sebagian return dengan _
//
// Penjelasan:
//   Pemanggil tidak wajib menangkap semua return. Blank identifier _
//   mengabaikan nilai di posisi tersebut. Contoh: v, _ := f() atau _, err := f().
//
// Alasan penggunaan:
//   Kadang hanya satu (atau beberapa) nilai yang dibutuhkan. _ menghindari
//   variabel tidak terpakai dan membuat intent "sengaja diabaikan" jelas.
//
// Fungsi penggunaan:
//   q, _ := divide(10, 3) hanya butuh hasil bagi. _, _, c := getTiga() hanya c.
//
// Output:
//   Hanya nilai pertama: 10
//   Hanya nilai ketiga: 3
// -----------------------------------------------------------------------------
func ignoreSebagianReturn() {
	first, _ := getPoint()
	fmt.Printf("  Hanya nilai pertama: %d\n", first)
	_, _, third := getTiga()
	fmt.Printf("  Hanya nilai ketiga: %d\n", third)
}

// -----------------------------------------------------------------------------
// 5. Named return untuk multiple value
//
// Penjelasan:
//   Tiap return value bisa diberi nama: func nama() (a int, b int). Variabel
//   a, b dianggap dideklarasikan di awal fungsi. return tanpa argumen
//   mengembalikan nilai current a dan b (naked return).
//
// Alasan penggunaan:
//   Signature jadi dokumentasi; nama menjelaskan arti tiap return. Naked
//   return mempersingkat kode di fungsi pendek.
//
// Fungsi penggunaan:
//   min, max := minMax(3, 7). Di dalam minMax: min, max = 3, 7; return.
//
// Output:
//   minMax(3, 7) = 3, 7
// -----------------------------------------------------------------------------
func namedMultipleReturn() {
	mn, mx := minMax(3, 7)
	fmt.Printf("  minMax(3, 7) = %d, %d\n", mn, mx)
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
// 6. Studi kasus: bagi dan sisa
//
// Penjelasan:
//   Fungsi divide(a, b int) (quotient, remainder int) mengembalikan hasil
//   bagi bulat dan sisa bagi dalam satu pemanggilan. Contoh klasik multiple return.
//
// Alasan penggunaan:
//   Pembagian integer dan modulo sering dipakai bersama; satu fungsi
//   mengembalikan keduanya tanpa dua pemanggilan atau struct.
//
// Fungsi penggunaan:
//   q, r := divide(17, 5) → q=3, r=2.
//
// Output:
//   divide(17, 5) = quotient 3, remainder 2
// -----------------------------------------------------------------------------
func kasusBagiSisa() {
	q, r := divide(17, 5)
	fmt.Printf("  divide(17, 5) = quotient %d, remainder %d\n", q, r)
}

func divide(a, b int) (int, int) {
	return a / b, a % b
}

// -----------------------------------------------------------------------------
// 7. Studi kasus: parse nama (depan, belakang)
//
// Penjelasan:
//   Fungsi splitNama(namaLengkap string) (depan, belakang string) memecah
//   string nama menjadi dua bagian (asumsi dipisah spasi). Multiple return
//   untuk dua string tanpa slice atau struct.
//
// Alasan penggunaan:
//   Form/API sering butuh nama depan dan belakang terpisah; satu return
//   memberikan keduanya sekaligus.
//
// Fungsi penggunaan:
//   depan, belakang := splitNama("Budi Santoso"). Jika satu kata, belakang "".
//
// Output:
//   splitNama("Budi Santoso") = depan: Budi, belakang: Santoso
// -----------------------------------------------------------------------------
func kasusParseNama() {
	depan, belakang := splitNama("Budi Santoso")
	fmt.Printf("  splitNama(\"Budi Santoso\") = depan: %s, belakang: %s\n", depan, belakang)
}

func splitNama(namaLengkap string) (depan, belakang string) {
	parts := strings.Fields(namaLengkap)
	if len(parts) >= 2 {
		return parts[0], parts[len(parts)-1]
	}
	if len(parts) == 1 {
		return parts[0], ""
	}
	return "", ""
}

// -----------------------------------------------------------------------------
// 8. Studi kasus: statistik slice (sum, avg)
//
// Penjelasan:
//   Fungsi statistik(nums []int) (sum int, avg float64) mengembalikan
//   jumlah dan rata-rata dari slice. Dua nilai terkait dari satu iterasi.
//
// Alasan penggunaan:
//   Satu loop menghitung sum sekaligus; avg = sum/len. Multiple return
//   menghindari dua fungsi terpisah atau struct kecil hanya untuk dua angka.
//
// Fungsi penggunaan:
//   sum, avg := statistik([]int{10, 20, 30}). sum=60, avg=20.
//
// Output:
//   statistik([10 20 30]) = sum 60, avg 20.00
// -----------------------------------------------------------------------------
func kasusStatistikSlice() {
	nums := []int{10, 20, 30}
	s, a := statistik(nums)
	fmt.Printf("  statistik(%v) = sum %d, avg %.2f\n", nums, s, a)
}

func statistik(nums []int) (sum int, avg float64) {
	if len(nums) == 0 {
		return 0, 0
	}
	for _, n := range nums {
		sum += n
	}
	avg = float64(sum) / float64(len(nums))
	return sum, avg
}

// -----------------------------------------------------------------------------
// 9. Studi kasus: lookup map (value, ok)
//
// Penjelasan:
//   Pola lookup map: v, ok := m[key]. ok true jika key ada, false jika tidak.
//   Ini bentuk multiple return: nilai + boolean "found". Bisa dibungkus
//   dalam fungsi yang return (value T, found bool).
//
// Alasan penggunaan:
//   Membedakan "key tidak ada" (ok=false) dengan "value zero" (ok=true, v=0).
//   Tanpa ok, v=0 bisa berarti tidak ada atau nilai memang 0.
//
// Fungsi penggunaan:
//   val, found := lookup(phonebook, "Budi"). if found { pakai val }.
//
// Output:
//   lookup("Budi") = 08123456789 true
//   lookup("Ani") =  false
// -----------------------------------------------------------------------------
func kasusLookupMap() {
	phonebook := map[string]string{"Budi": "08123456789", "Siti": "08234567890"}
	val, found := lookup(phonebook, "Budi")
	fmt.Printf("  lookup(\"Budi\") = %s %t\n", val, found)
	val2, found2 := lookup(phonebook, "Ani")
	fmt.Printf("  lookup(\"Ani\") = %s %t\n", val2, found2)
}

func lookup(m map[string]string, key string) (string, bool) {
	v, ok := m[key]
	return v, ok
}
