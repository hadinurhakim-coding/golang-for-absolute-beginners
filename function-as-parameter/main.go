// =============================================================================
// FUNCTION AS PARAMETER DI GO
// =============================================================================
// Fungsi bisa diteruskan sebagai argument ke fungsi lain. Parameter bertipe
// fungsi: func(T) R atau func(T1, T2) R. Pemanggil mengirim nama fungsi atau
// literal. Pola umum: callback, mapper, predicate, comparator.
//
// POLA KOMENTAR (konsisten):
//   1. Penjelasan        — ringkasan materi / apa yang dilakukan.
//   2. Alasan penggunaan — kenapa fitur ini dipakai.
//   3. Fungsi penggunaan — cara pakai (pemanggil, signature parameter).
//   4. Output            — contoh hasil cetakan/return dari fungsi tersebut.
//   5. Penjelasan output — mengapa output-nya bisa seperti itu (alur/logika).
// =============================================================================

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("========== 1. Dasar: parameter bertipe fungsi ==========")
	dasarParameterFungsi()

	fmt.Println("\n========== 2. Callback: fungsi dipanggil di dalam ==========")
	callbackDasar()

	fmt.Println("\n========== 3. Fungsi dengan banyak parameter (termasuk fungsi) ==========")
	parameterBanyak()

	fmt.Println("\n========== 4. Anonymous function sebagai argument ==========")
	anonymousAsArg()

	fmt.Println("\n========== 5. Studi kasus: map slice (transform tiap elemen) ==========")
	kasusMapSlice()

	fmt.Println("\n========== 6. Studi kasus: filter slice (predikat) ==========")
	kasusFilterSlice()

	fmt.Println("\n========== 7. Studi kasus: forEach (efek samping per elemen) ==========")
	kasusForEach()

	fmt.Println("\n========== 8. Studi kasus: string transform (uppercase/lowercase) ==========")
	kasusStringTransform()
}

// -----------------------------------------------------------------------------
// 1. Dasar: parameter bertipe fungsi
//
// Penjelasan:
//   Parameter fn bertipe func(int) int. Pemanggil wajib mengirim fungsi yang
//   menerima int dan mengembalikan int. Di dalam body, fn dipanggil seperti
//   fungsi biasa: fn(n).
//
// Alasan penggunaan:
//   Memisahkan "alur" (berapa kali, untuk nilai apa) dari "perilaku" (apa
//   yang dilakukan). Satu fungsi bisa dipakai dengan banyak implementasi.
//
// Fungsi penggunaan:
//   apply(5, double) — kirim fungsi double; apply memanggil double(5) dan return hasilnya.
//
// Output:
//   apply(5, double) = 10
//
// Penjelasan output:
//   apply menerima n=5 dan fn=double; di dalam apply, fn(5) memanggil double(5)
//   yang return 10, sehingga apply return 10 dan cetakan "10".
// -----------------------------------------------------------------------------
func dasarParameterFungsi() {
	fmt.Printf("  apply(5, double) = %d\n", apply(5, double))
}

func apply(n int, fn func(int) int) int {
	return fn(n)
}

func double(x int) int {
	return x * 2
}

// -----------------------------------------------------------------------------
// 2. Callback: fungsi dipanggil di dalam
//
// Penjelasan:
//   "Callback" = fungsi yang dikirim untuk dipanggil nanti oleh fungsi lain.
//   repeat(3, printer) memanggil printer() sebanyak 3 kali. printer adalah
//   parameter fungsi (tanpa argument, tanpa return).
//
// Alasan penggunaan:
//   Mengabstraksi "apa yang dijalankan" dari "berapa kali / kapan". Berguna
//   untuk event handler, retry, iterator.
//
// Fungsi penggunaan:
//   repeat(3, printer) — di dalam repeat, loop 3 kali dan setiap kali panggil printer().
//
// Output:
//   repeat(3, printer): Hello Hello Hello
//
// Penjelasan output:
//   repeat memanggil fn() tiga kali; fn = printer yang cetak "Hello ", sehingga
//   tiga kali cetak "Hello " lalu newline.
// -----------------------------------------------------------------------------
func callbackDasar() {
	fmt.Print("  repeat(3, printer): ")
	repeat(3, printer)
	fmt.Println()
}

func repeat(n int, fn func()) {
	for i := 0; i < n; i++ {
		fn()
	}
}

func printer() {
	fmt.Print("Hello ")
}

// -----------------------------------------------------------------------------
// 3. Fungsi dengan banyak parameter (termasuk fungsi)
//
// Penjelasan:
//   Fungsi boleh punya parameter biasa dan parameter bertipe fungsi. Contoh:
//   compute(a, b int, op func(int,int) int). op dipanggil dengan a dan b.
//
// Alasan penggunaan:
//   Satu fungsi "compute" dipakai untuk berbagai operasi (add, minus, multiply)
//   tanpa menulis computeAdd, computeMinus, dll. Operasi dikirim sebagai parameter.
//
// Fungsi penggunaan:
//   compute(10, 3, add) = 13; compute(10, 3, minus) = 7.
//
// Output:
//   compute(10, 3, add) = 13
//   compute(10, 3, minus) = 7
//
// Penjelasan output:
//   compute(10, 3, add) memanggil op(10,3) dengan op=add, return 13. Same for
//   minus: op(10,3) = minus(10,3) = 7.
// -----------------------------------------------------------------------------
func parameterBanyak() {
	fmt.Printf("  compute(10, 3, add) = %d\n", compute(10, 3, add))
	fmt.Printf("  compute(10, 3, minus) = %d\n", compute(10, 3, minus))
}

func compute(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

// -----------------------------------------------------------------------------
// 4. Anonymous function sebagai argument
//
// Penjelasan:
//   Argument parameter fungsi bisa berupa literal (anonymous function), tidak
//   harus nama fungsi. Langsung tulis func(x int) int { return x * x } sebagai
//   argument. Berguna untuk logic sekali pakai.
//
// Alasan penggunaan:
//   Tidak perlu mendefinisikan fungsi terpisah jika hanya dipakai sekali;
//   kode lebih ringkas dan dekat dengan pemanggilan.
//
// Fungsi penggunaan:
//   apply(4, func(x int) int { return x * x }) = 16. Literal jadi parameter fn.
//
// Output:
//   apply(4, func(x int) int { return x*x }) = 16
//
// Penjelasan output:
//   apply(4, ...) memanggil fn(4) dengan fn = fungsi literal yang return x*x;
//   x=4 maka 4*4=16, sehingga output "16".
// -----------------------------------------------------------------------------
func anonymousAsArg() {
	square := func(x int) int { return x * x }
	fmt.Printf("  apply(4, func(x int) int { return x*x }) = %d\n", apply(4, square))
}

// -----------------------------------------------------------------------------
// 5. Studi kasus: map slice (transform tiap elemen)
//
// Penjelasan:
//   mapInts(slice []int, fn func(int) int) mengembalikan slice baru: tiap
//   elemen e diganti jadi fn(e). Parameter fn adalah "mapper".
//
// Alasan penggunaan:
//   Satu fungsi mapInts dipakai untuk berbagai transformasi (double, square,
//   negate) dengan mengirim fungsi yang berbeda. Reusable higher-order function.
//
// Fungsi penggunaan:
//   mapInts([]int{1,2,3}, double) = [2 4 6]. Tiap elemen di-double.
//
// Output:
//   mapInts([1 2 3], double) = [2 4 6]
//
// Penjelasan output:
//   Untuk setiap 1, 2, 3 dipanggil double(e): 2, 4, 6. Hasil append ke out
//   sehingga slice [2 4 6] dicetak.
// -----------------------------------------------------------------------------
func kasusMapSlice() {
	nums := []int{1, 2, 3}
	out := mapInts(nums, double)
	fmt.Printf("  mapInts([1 2 3], double) = %v\n", out)
}

func mapInts(slice []int, fn func(int) int) []int {
	out := make([]int, len(slice))
	for i, v := range slice {
		out[i] = fn(v)
	}
	return out
}

// -----------------------------------------------------------------------------
// 6. Studi kasus: filter slice (predikat)
//
// Penjelasan:
//   filterInts(slice []int, pred func(int) bool) mengembalikan slice berisi
//   hanya elemen e yang pred(e) true. Parameter pred adalah predicate function.
//
// Alasan penggunaan:
//   Satu filterInts dipakai untuk genap, ganjil, positif, > 5, dll dengan
//   mengirim fungsi predikat yang berbeda.
//
// Fungsi penggunaan:
//   filterInts([]int{1,2,3,4,5}, genap) = [2 4].
//
// Output:
//   filterInts([1 2 3 4 5], genap) = [2 4]
//
// Penjelasan output:
//   pred = genap; hanya 2 dan 4 yang genap (genap(2)=true, genap(4)=true);
//   elemen lain tidak di-append, sehingga hasil [2 4].
// -----------------------------------------------------------------------------
func kasusFilterSlice() {
	nums := []int{1, 2, 3, 4, 5}
	out := filterInts(nums, genap)
	fmt.Printf("  filterInts([1 2 3 4 5], genap) = %v\n", out)
}

func filterInts(slice []int, pred func(int) bool) []int {
	var out []int
	for _, v := range slice {
		if pred(v) {
			out = append(out, v)
		}
	}
	return out
}

func genap(n int) bool {
	return n%2 == 0
}

// -----------------------------------------------------------------------------
// 7. Studi kasus: forEach (efek samping per elemen)
//
// Penjelasan:
//   forEachInts(slice []int, fn func(int)) memanggil fn(e) untuk tiap elemen e.
//   Tidak mengembalikan nilai; dipakai untuk side effect (print, log, kirim).
//
// Alasan penggunaan:
//   Iterasi dengan perilaku yang bisa diubah dari luar. Satu forEach dipakai
//   untuk print, sum ke variabel luar, atau kirim ke channel, dengan mengirim
//   fn yang berbeda.
//
// Fungsi penggunaan:
//   forEachInts([]int{1,2,3}, func(n int) { fmt.Print(n, " ") }) cetak "1 2 3 ".
//
// Output:
//   forEachInts([1 2 3], print): 1 2 3
//
// Penjelasan output:
//   fn dipanggil untuk 1, 2, 3; fn = printerInt yang cetak nilai + spasi,
//   sehingga output "1 2 3 ".
// -----------------------------------------------------------------------------
func kasusForEach() {
	fmt.Print("  forEachInts([1 2 3], print): ")
	forEachInts([]int{1, 2, 3}, printerInt)
	fmt.Println()
}

func forEachInts(slice []int, fn func(int)) {
	for _, v := range slice {
		fn(v)
	}
}

func printerInt(n int) {
	fmt.Print(n, " ")
}

// -----------------------------------------------------------------------------
// 8. Studi kasus: string transform (uppercase/lowercase)
//
// Penjelasan:
//   transformString(s string, fn func(string) string) mengembalikan fn(s).
//   Pemanggil mengirim strings.ToUpper atau strings.ToLower (atau literal)
//   sebagai parameter. Menggambarkan fungsi dari paket lain sebagai parameter.
//
// Alasan penggunaan:
//   Satu fungsi "transform" dipakai untuk berbagai aturan (upper, lower, trim,
//   replace) tanpa menulis transformUpper, transformLower, dll.
//
// Fungsi penggunaan:
//   transformString("Hello", strings.ToUpper) = "HELLO".
//
// Output:
//   transformString("Hello", ToUpper) = HELLO
//   transformString("Hello", ToLower) = hello
//
// Penjelasan output:
//   fn("Hello") dengan fn=strings.ToUpper mengembalikan "HELLO"; dengan
//   fn=strings.ToLower mengembalikan "hello". Keduanya dicetak.
// -----------------------------------------------------------------------------
func kasusStringTransform() {
	fmt.Printf("  transformString(\"Hello\", ToUpper) = %s\n", transformString("Hello", strings.ToUpper))
	fmt.Printf("  transformString(\"Hello\", ToLower) = %s\n", transformString("Hello", strings.ToLower))
}

func transformString(s string, fn func(string) string) string {
	return fn(s)
}
