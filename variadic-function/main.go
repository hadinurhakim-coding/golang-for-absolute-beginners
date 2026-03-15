// =============================================================================
// VARIADIC FUNCTION DI GO
// =============================================================================
// Variadic function = fungsi yang menerima nol atau lebih argument dengan tipe
// yang sama. Bentuk parameter: nama ...T. Di dalam fungsi, nama berperilaku
// seperti slice []T. Hanya satu parameter variadic per fungsi dan harus di
// posisi terakhir. Pemanggil: f(1,2,3) atau f(slice...).
//
// POLA KOMENTAR (konsisten):
//   1. Penjelasan        — ringkasan materi / apa yang dilakukan.
//   2. Alasan penggunaan — kenapa fitur ini dipakai.
//   3. Fungsi penggunaan — cara pakai (pemanggil, spread slice).
//   4. Output            — contoh hasil cetakan/return dari fungsi tersebut.
//   5. Penjelasan output — mengapa output-nya bisa seperti itu (alur/logika).
// =============================================================================

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("========== 1. Dasar: parameter ...T ==========")
	dasarVariadic()

	fmt.Println("\n========== 2. Nol argument (variadic boleh kosong) ==========")
	variadicNolArg()

	fmt.Println("\n========== 3. Spread slice dengan ... ==========")
	spreadSlice()

	fmt.Println("\n========== 4. Variadic dengan parameter biasa (variadic di akhir) ==========")
	variadicDenganParamBiasa()

	fmt.Println("\n========== 5. Studi kasus: sum banyak angka ==========")
	kasusSum()

	fmt.Println("\n========== 6. Studi kasus: join string dengan pemisah ==========")
	kasusJoin()

	fmt.Println("\n========== 7. Studi kasus: nilai maksimum ==========")
	kasusMax()

	fmt.Println("\n========== 8. Studi kasus: gabung slice (append style) ==========")
	kasusGabungSlice()

	fmt.Println("\n========== 9. Studi kasus: cetak dengan prefix (variadic interface{}) ==========")
	kasusPrintPrefix()
}

// -----------------------------------------------------------------------------
// 1. Dasar: parameter ...T
//
// Penjelasan:
//   Parameter ...int menerima nol atau lebih int. Di dalam fungsi, nums
//   bertipe []int. Iterasi dengan for _, n := range nums.
//
// Alasan penggunaan:
//   Jumlah argument tidak tetap tanpa harus membuat slice dulu. Pemanggil
//   bisa tulis sum(1,2,3) atau sum(1,2,3,4,5) tanpa bungkus dalam slice.
//
// Fungsi penggunaan:
//   total := sumInts(1, 2, 3). Di dalam: for range nums.
//
// Output:
//   sumInts(1, 2, 3) = 6
//
// Penjelasan output:
//   Argument 1, 2, 3 dikumpulkan jadi slice nums; loop menjumlahkan 1+2+3 = 6,
//   lalu return 6 sehingga cetakan menjadi "6".
// -----------------------------------------------------------------------------
func dasarVariadic() {
	fmt.Printf("  sumInts(1, 2, 3) = %d\n", sumInts(1, 2, 3))
}

func sumInts(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// -----------------------------------------------------------------------------
// 2. Nol argument (variadic boleh kosong)
//
// Penjelasan:
//   Variadic parameter boleh menerima nol argument. Pemanggil: f(). Di dalam
//   fungsi slice akan len(nums)==0. Perlu handle agar tidak panic (misalnya
//   return 0 atau nilai default).
//
// Alasan penggunaan:
//   Fleksibel: kadang pemanggil tidak punya nilai untuk dikirim. Fungsi
//   harus defensif (cek len sebelum akses indeks).
//
// Fungsi penggunaan:
//   sumInts() → 0. Di dalam: if len(nums)==0 { return 0 }.
//
// Output:
//   sumInts() = 0
//
// Penjelasan output:
//   Tanpa argument, nums berisi slice kosong; loop tidak jalan, total tetap 0,
//   sehingga return 0 dan cetakan menampilkan "0".
// -----------------------------------------------------------------------------
func variadicNolArg() {
	fmt.Printf("  sumInts() = %d\n", sumInts())
}

// -----------------------------------------------------------------------------
// 3. Spread slice dengan ...
//
// Penjelasan:
//   Jika nilai sudah ada di slice, bisa "melebarkan" jadi argument variadic
//   dengan slice... . Contoh: nums := []int{1,2,3}; sumInts(nums...).
//
// Alasan penggunaan:
//   Menghindari konversi manual atau loop; slice langsung dipakai sebagai
//   sumber argument untuk fungsi variadic.
//
// Fungsi penggunaan:
//   s := []int{10, 20, 30}; sumInts(s...) = 60.
//
// Output:
//   sumInts([]int{10, 20, 30}...) = 60
//
// Penjelasan output:
//   s... melebarkan slice [10,20,30] jadi tiga argument 10, 20, 30; sumInts
//   menjumlahkannya (10+20+30=60) sehingga output "60".
// -----------------------------------------------------------------------------
func spreadSlice() {
	s := []int{10, 20, 30}
	fmt.Printf("  sumInts([]int{10, 20, 30}...) = %d\n", sumInts(s...))
}

// -----------------------------------------------------------------------------
// 4. Variadic dengan parameter biasa (variadic di akhir)
//
// Penjelasan:
//   Fungsi boleh punya parameter biasa lalu satu variadic di posisi terakhir.
//   Bentuk: func nama(prefix string, values ...int). Pemanggil: nama("x", 1, 2, 3).
//   Parameter biasa tidak boleh setelah variadic.
//
// Alasan penggunaan:
//   Satu atau beberapa nilai "wajib" (prefix, separator) plus banyak nilai
//   opsional (list angka, list string).
//
// Fungsi penggunaan:
//   joinWithSep(", ", "a", "b", "c") → "a, b, c".
//
// Output:
//   joinWithSep(", ", "a", "b", "c") = a, b, c
//
// Penjelasan output:
//   sep = ", ", parts = ["a","b","c"]; strings.Join(parts, sep) menggabungkan
//   dengan ", " di antara tiap elemen sehingga hasil "a, b, c".
// -----------------------------------------------------------------------------
func variadicDenganParamBiasa() {
	fmt.Printf("  joinWithSep(\", \", \"a\", \"b\", \"c\") = %s\n", joinWithSep(", ", "a", "b", "c"))
}

func joinWithSep(sep string, parts ...string) string {
	return strings.Join(parts, sep)
}

// -----------------------------------------------------------------------------
// 5. Studi kasus: sum banyak angka
//
// Penjelasan:
//   sum(nums ...int) menjumlahkan semua argument. Menggambarkan variadic
//   paling sederhana: iterasi slice dan agregasi.
//
// Alasan penggunaan:
//   Jumlah angka bisa berubah; variadic memudahkan pemanggilan tanpa
//   membuat slice terlebih dahulu.
//
// Fungsi penggunaan:
//   sum(1, 2, 3, 4, 5) = 15. Atau sum(slice...).
//
// Output:
//   sum(1, 2, 3, 4, 5) = 15
//
// Penjelasan output:
//   Lima argument 1..5 dijumlahkan dalam loop: 1+2+3+4+5 = 15, sehingga
//   return 15 dan cetakan menampilkan "15".
// -----------------------------------------------------------------------------
func kasusSum() {
	fmt.Printf("  sum(1, 2, 3, 4, 5) = %d\n", sum(1, 2, 3, 4, 5))
}

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// -----------------------------------------------------------------------------
// 6. Studi kasus: join string dengan pemisah
//
// Penjelasan:
//   join(sep string, parts ...string) menggabungkan semua parts dengan sep
//   di antara. Parameter pertama pemisah (biasa), sisanya variadic string.
//
// Alasan penggunaan:
//   Mirip strings.Join tapi dengan variadic; pemanggil tidak perlu bikin
//   slice: join(" - ", "A", "B", "C").
//
// Fungsi penggunaan:
//   join(" - ", "A", "B", "C") = "A - B - C".
//
// Output:
//   join(" - ", "A", "B", "C") = A - B - C
//
// Penjelasan output:
//   parts = ["A","B","C"], sep = " - "; strings.Join menempatkan " - " di
//   antara A, B, dan C sehingga string hasil "A - B - C".
// -----------------------------------------------------------------------------
func kasusJoin() {
	fmt.Printf("  join(\" - \", \"A\", \"B\", \"C\") = %s\n", join(" - ", "A", "B", "C"))
}

func join(sep string, parts ...string) string {
	return strings.Join(parts, sep)
}

// -----------------------------------------------------------------------------
// 7. Studi kasus: nilai maksimum
//
// Penjelasan:
//   max(nums ...int) mengembalikan nilai terbesar. Jika len(nums)==0 bisa
//   return 0 atau nilai sentinel/error. Di sini return 0 untuk kosong.
//
// Alasan penggunaan:
//   Mencari max dari N bilangan tanpa harus membuat slice; variadic
//   memudahkan: max(3, 7, 2, 9).
//
// Fungsi penggunaan:
//   max(3, 7, 2, 9) = 9.
//
// Output:
//   max(3, 7, 2, 9) = 9
//
// Penjelasan output:
//   nums = [3,7,2,9]; m mulai 3, lalu dibandingkan dengan 7 (m=7), 2 (tetap 7),
//   9 (m=9). Nilai terbesar 9 dikembalikan sehingga output "9".
// -----------------------------------------------------------------------------
func kasusMax() {
	fmt.Printf("  max(3, 7, 2, 9) = %d\n", max(3, 7, 2, 9))
}

func max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	m := nums[0]
	for _, n := range nums[1:] {
		if n > m {
			m = n
		}
	}
	return m
}

// -----------------------------------------------------------------------------
// 8. Studi kasus: gabung slice (append style)
//
// Penjelasan:
//   concat(slices ...[]int) menggabungkan semua slice menjadi satu. Di dalam
//   loop range, append ke hasil. Menggambarkan variadic dengan tipe slice.
//
// Alasan penggunaan:
//   Menggabungkan beberapa slice tanpa menulis append(append(s1, s2...), s3...).
//   Satu pemanggil: concat(s1, s2, s3).
//
// Fungsi penggunaan:
//   concat([]int{1,2}, []int{3,4}) = [1 2 3 4].
//
// Output:
//   concat([1 2], [3 4]) = [1 2 3 4]
//
// Penjelasan output:
//   slices = [[1,2], [3,4]]; loop append out dengan [1,2] lalu [3,4] (pakai s...),
//   sehingga out = [1,2,3,4] dan %v mencetak "[1 2 3 4]".
// -----------------------------------------------------------------------------
func kasusGabungSlice() {
	a := []int{1, 2}
	b := []int{3, 4}
	fmt.Printf("  concat([1 2], [3 4]) = %v\n", concat(a, b))
}

func concat(slices ...[]int) []int {
	var out []int
	for _, s := range slices {
		out = append(out, s...)
	}
	return out
}

// -----------------------------------------------------------------------------
// 9. Studi kasus: cetak dengan prefix (variadic interface{})
//
// Penjelasan:
//   printPrefix(prefix string, args ...interface{}) mencetak prefix lalu
//   nilai-nilai args (mirip fmt.Println). ...interface{} menerima tipe apa saja.
//
// Alasan penggunaan:
//   Fungsi log/debug yang menerima banyak nilai bertipe bebas; interface{}
//   membuat variadic fleksibel untuk print.
//
// Fungsi penggunaan:
//   printPrefix("[INFO]", "nilai", 42, true). Output: [INFO] nilai 42 true
//
// Output:
//   [INFO] nilai 42 true
//
// Penjelasan output:
//   prefix = "[INFO]", args = ["nilai", 42, true]; fmt.Print(prefix, " ") cetak
//   "[INFO] ", lalu fmt.Println(args...) cetak nilai 42 true dengan spasi, sehingga
//   satu baris: "[INFO] nilai 42 true".
// -----------------------------------------------------------------------------
func kasusPrintPrefix() {
	fmt.Print("  ")
	printPrefix("[INFO]", "nilai", 42, true)
	fmt.Println()
}

func printPrefix(prefix string, args ...interface{}) {
	fmt.Print(prefix, " ")
	fmt.Println(args...)
}
