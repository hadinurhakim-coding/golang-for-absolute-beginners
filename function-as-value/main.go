// =============================================================================
// FUNCTION AS VALUE DI GO
// =============================================================================
// Di Go, fungsi adalah "first-class": bisa disimpan di variabel, dikirim sebagai
// argument, dan dikembalikan dari fungsi lain. Tipe fungsi: func(param T) R atau
// func(T) R. Anonymous function: func() { } atau func(x int) int { return x*2 }.
//
// POLA KOMENTAR (konsisten):
//   1. Penjelasan        — ringkasan materi / apa yang dilakukan.
//   2. Alasan penggunaan — kenapa fitur ini dipakai.
//   3. Fungsi penggunaan  — cara pakai (assign, pass, return).
//   4. Output            — contoh hasil cetakan/return dari fungsi tersebut.
//   5. Penjelasan output — mengapa output-nya bisa seperti itu (alur/logika).
// =============================================================================

package main

import "fmt"

func main() {
	fmt.Println("========== 1. Fungsi disimpan di variabel ==========")
	fungsiDiVariabel()

	fmt.Println("\n========== 2. Tipe fungsi (func(T) R) ==========")
	tipeFungsi()

	fmt.Println("\n========== 3. Fungsi sebagai parameter (callback) ==========")
	fungsiSebagaiParameter()

	fmt.Println("\n========== 4. Anonymous function (literal) ==========")
	anonymousFunction()

	fmt.Println("\n========== 5. Fungsi mengembalikan fungsi ==========")
	fungsiReturnFungsi()

	fmt.Println("\n========== 6. Studi kasus: apply (operasi ke nilai) ==========")
	kasusApply()

	fmt.Println("\n========== 7. Studi kasus: kalkulator (operasi sebagai value) ==========")
	kasusKalkulator()

	fmt.Println("\n========== 8. Studi kasus: filter slice dengan fungsi ==========")
	kasusFilterSlice()
}

// -----------------------------------------------------------------------------
// 1. Fungsi disimpan di variabel
//
// Penjelasan:
//   Fungsi bisa di-assign ke variabel. Variabel tersebut punya tipe fungsi
//   dan bisa dipanggil seperti fungsi: f(). Nama fungsi asli tidak wajib dipakai.
//
// Alasan penggunaan:
//   Memilih implementasi di runtime (strategi, plugin), atau menyimpan referensi
//   untuk dipanggil nanti (callback, handler).
//
// Fungsi penggunaan:
//   var f func(int) int = double; f(5) memanggil double(5).
//
// Output:
//   double(5) = 10
//   lewat variabel f(5) = 10
//
// Penjelasan output:
//   f menunjuk ke fungsi double; f(5) memanggil double(5) yang return 10,
//   sehingga kedua baris mencetak 10.
// -----------------------------------------------------------------------------
func fungsiDiVariabel() {
	fmt.Printf("  double(5) = %d\n", double(5))
	f := double
	fmt.Printf("  lewat variabel f(5) = %d\n", f(5))
}

func double(x int) int {
	return x * 2
}

// -----------------------------------------------------------------------------
// 2. Tipe fungsi (func(T) R)
//
// Penjelasan:
//   Tipe fungsi ditulis func(parameter) return. Contoh: func(int) int untuk
//   fungsi yang menerima int dan mengembalikan int. Bisa dipakai untuk variabel,
//   parameter, atau field struct.
//
// Alasan penggunaan:
//   Agar variabel/parameter hanya menerima fungsi dengan signature tertentu;
//   type safety dan dokumentasi di signature.
//
// Fungsi penggunaan:
//   var op func(int, int) int = add; op(2,3)=5. Type alias: type BinOp func(int,int) int.
//
// Output:
//   add(2, 3) = 5
//   op(2, 3) = 5
//
// Penjelasan output:
//   op bertipe func(int,int) int dan di-assign add; op(2,3) memanggil add(2,3)
//   yang return 5, sehingga output "5".
// -----------------------------------------------------------------------------
func tipeFungsi() {
	fmt.Printf("  add(2, 3) = %d\n", add(2, 3))
	var op func(int, int) int = add
	fmt.Printf("  op(2, 3) = %d\n", op(2, 3))
}

func add(a, b int) int {
	return a + b
}

// -----------------------------------------------------------------------------
// 3. Fungsi sebagai parameter (callback)
//
// Penjelasan:
//   Parameter bisa bertipe fungsi. Pemanggil mengirim fungsi (atau variabel
//   yang berisi fungsi); di dalam fungsi dipanggil. Pola umum: callback, mapper.
//
// Alasan penggunaan:
//   Memisahkan "apa yang dilakukan" (logic) dari "kapan/berapa kali" (loop, API).
//   Contoh: for each element call fn(v).
//
// Fungsi penggunaan:
//   apply(10, double) memanggil double(10) di dalam apply dan return 20.
//
// Output:
//   apply(10, double) = 20
//
// Penjelasan output:
//   apply menerima 10 dan fungsi double; di dalam apply, fn(10) memanggil
//   double(10) yang return 20, sehingga output "20".
// -----------------------------------------------------------------------------
func fungsiSebagaiParameter() {
	fmt.Printf("  apply(10, double) = %d\n", apply(10, double))
}

func apply(n int, fn func(int) int) int {
	return fn(n)
}

// -----------------------------------------------------------------------------
// 4. Anonymous function (literal)
//
// Penjelasan:
//   Fungsi tanpa nama bisa didefinisikan inline: func(x int) int { return x + 1 }.
//   Bisa di-assign ke variabel atau dipanggil langsung: func() { fmt.Println("hi") }().
//   Tanda () di akhir memanggil fungsi tersebut.
//
// Alasan penggunaan:
//   Untuk callback sekali pakai, atau closure yang menangkap variabel di sekitarnya,
//   tanpa perlu mendefinisikan fungsi terpisah.
//
// Fungsi penggunaan:
//   inc := func(x int) int { return x+1 }; inc(5)=6. Atau langsung: func(){}().
//
// Output:
//   anonymous inc(5) = 6
//   IIFE: Hello
//
// Penjelasan output:
//   inc adalah fungsi literal yang return x+1; inc(5) return 6. IIFE (Immediately
//   Invoked Function Expression) func(){ fmt.Println("Hello") }() langsung jalan
//   sehingga mencetak "Hello".
// -----------------------------------------------------------------------------
func anonymousFunction() {
	inc := func(x int) int { return x + 1 }
	fmt.Printf("  anonymous inc(5) = %d\n", inc(5))
	fmt.Print("  IIFE: ")
	func() { fmt.Println("Hello") }()
}

// -----------------------------------------------------------------------------
// 5. Fungsi mengembalikan fungsi
//
// Penjelasan:
//   Return type bisa tipe fungsi: func() func(int) int. Fungsi yang dikembalikan
//   bisa "mengingat" variabel di scope pembuatnya (closure). Contoh: multiplier(k)
//   return fungsi yang mengalikan dengan k.
//
// Alasan penggunaan:
//   Factory: membuat fungsi dengan konfigurasi tertentu. Closure menyimpan state
//   (misalnya multiplier(3) menghasilkan fungsi "kali 3").
//
// Fungsi penggunaan:
//   times3 := multiplier(3); times3(4)=12. Fungsi return mengembalikan func(int) int.
//
// Output:
//   multiplier(3)(4) = 12
//   times3(5) = 15
//
// Penjelasan output:
//   multiplier(3) mengembalikan fungsi yang return n*3; pemanggilan dengan 4
//   memberi 12. times3 menyimpan fungsi itu; times3(5) memanggil n*3 dengan n=5
//   sehingga 15.
// -----------------------------------------------------------------------------
func fungsiReturnFungsi() {
	fmt.Printf("  multiplier(3)(4) = %d\n", multiplier(3)(4))
	times3 := multiplier(3)
	fmt.Printf("  times3(5) = %d\n", times3(5))
}

func multiplier(k int) func(int) int {
	return func(n int) int {
		return n * k
	}
}

// -----------------------------------------------------------------------------
// 6. Studi kasus: apply (operasi ke nilai)
//
// Penjelasan:
//   applyOp(n int, op func(int) int) menerima nilai dan fungsi operasi, lalu
//   mengembalikan op(n). Pemanggil bisa mengirim double, square, atau fungsi
//   literal.
//
// Alasan penggunaan:
//   Satu fungsi "apply" dipakai untuk berbagai transformasi tanpa menulis
//   applyDouble, applySquare, dll. Pola higher-order function.
//
// Fungsi penggunaan:
//   applyOp(5, square) = 25. square(5) return 5*5.
//
// Output:
//   applyOp(5, square) = 25
//
// Penjelasan output:
//   applyOp memanggil op(5) dengan op = square; square(5) return 25, sehingga
//   output "25".
// -----------------------------------------------------------------------------
func kasusApply() {
	fmt.Printf("  applyOp(5, square) = %d\n", applyOp(5, square))
}

func square(n int) int {
	return n * n
}

func applyOp(n int, op func(int) int) int {
	return op(n)
}

// -----------------------------------------------------------------------------
// 7. Studi kasus: kalkulator (operasi sebagai value)
//
// Penjelasan:
//   Variabel calc bertipe func(int,int) int. Bisa di-assign add atau minus;
//   calc(a,b) memanggil operasi yang sedang disimpan. Menggambarkan strategi
//   lewat "function as value".
//
// Alasan penggunaan:
//   Memilih operasi di runtime tanpa banyak if/switch. Pola strategy.
//
// Fungsi penggunaan:
//   calc = add; calc(10,3)=13. Lalu calc = minus; calc(10,3)=7.
//
// Output:
//   calc = add: 10 + 3 = 13
//   calc = minus: 10 - 3 = 7
//
// Penjelasan output:
//   Pertama calc menunjuk add, calc(10,3) memanggil add(10,3)=13. Lalu calc
//   di-assign minus, calc(10,3) memanggil minus(10,3)=7.
// -----------------------------------------------------------------------------
func kasusKalkulator() {
	var calc func(int, int) int
	calc = add
	fmt.Printf("  calc = add: 10 + 3 = %d\n", calc(10, 3))
	calc = minus
	fmt.Printf("  calc = minus: 10 - 3 = %d\n", calc(10, 3))
}

func minus(a, b int) int {
	return a - b
}

// -----------------------------------------------------------------------------
// 8. Studi kasus: filter slice dengan fungsi
//
// Penjelasan:
//   filterInts(slice []int, pred func(int) bool) mengembalikan slice baru
//   berisi elemen yang pred(e) true. pred adalah "function as value" (predikat).
//
// Alasan penggunaan:
//   Satu fungsi filter dipakai untuk berbagai kriteria (genap, positif, > 5)
//   dengan mengirim fungsi predikat yang berbeda.
//
// Fungsi penggunaan:
//   filterInts([]int{1,2,3,4,5}, genap) = [2 4]. genap(e) return e%2==0.
//
// Output:
//   filterInts(..., genap) = [2 4]
//
// Penjelasan output:
//   pred = genap; untuk tiap elemen 1,2,3,4,5 hanya 2 dan 4 yang genap sehingga
//   genap(e) true; hasil filter [2 4] dicetak.
// -----------------------------------------------------------------------------
func kasusFilterSlice() {
	nums := []int{1, 2, 3, 4, 5}
	out := filterInts(nums, genap)
	fmt.Printf("  filterInts(..., genap) = %v\n", out)
}

func genap(n int) bool {
	return n%2 == 0
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
