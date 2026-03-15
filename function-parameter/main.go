// =============================================================================
// FUNCTION PARAMETER DI GO
// =============================================================================
// Parameter = nilai yang diterima fungsi saat dipanggil. Ditulis dalam ( ) setelah
// nama fungsi: func nama(param1 T1, param2 T2). Di Go, argument disalin ke
// parameter (pass by value) kecuali tipe reference (slice, map, channel) atau pointer.
//
// POLA KOMENTAR (konsisten):
//   1. Penjelasan   — ringkasan materi / apa yang dilakukan.
//   2. Alasan penggunaan — kenapa fitur ini dipakai.
//   3. Fungsi penggunaan — cara pakai (pemanggilan, parameter, return).
//   4. Output       — contoh hasil cetakan/return dari fungsi tersebut.
// =============================================================================

package main

import "fmt"

func main() {
	fmt.Println("========== 1. Satu parameter ==========")
	satuParameter()

	fmt.Println("\n========== 2. Banyak parameter (tipe sama & beda) ==========")
	banyakParameter()

	fmt.Println("\n========== 3. Parameter tipe sama disingkat ==========")
	parameterTipeSama()

	fmt.Println("\n========== 4. Pass by value (nilai disalin) ==========")
	passByValue()

	fmt.Println("\n========== 5. Parameter pointer (*T) ==========")
	parameterPointer()

	fmt.Println("\n========== 6. Variadic parameter (...T) ==========")
	parameterVariadic()

	fmt.Println("\n========== 7. Studi kasus: hitung total belanja ==========")
	kasusTotalBelanja()

	fmt.Println("\n========== 8. Studi kasus: tukar nilai dengan pointer ==========")
	kasusTukarNilai()

	fmt.Println("\n========== 9. Studi kasus: rata-rata banyak nilai (variadic) ==========")
	kasusRataRataVariadic()
}

// -----------------------------------------------------------------------------
// 1. Satu parameter
//
// Penjelasan:
//   Fungsi menerima tepat satu argument. Bentuk: func nama(param T). Saat
//   dipanggil nama(nilai), nilai disalin ke param.
//
// Alasan penggunaan:
//   Satu input cukup untuk operasi sederhana (cetak, validasi, transformasi
//   satu nilai).
//
// Fungsi penggunaan:
//   greet("Budi"). Parameter nama menerima string "Budi".
//
// Output:
//   Halo, Budi!
// -----------------------------------------------------------------------------
func satuParameter() {
	greet("Budi")
}

func greet(nama string) {
	fmt.Printf("  Halo, %s!\n", nama)
}

// -----------------------------------------------------------------------------
// 2. Banyak parameter (tipe sama dan tipe beda)
//
// Penjelasan:
//   Beberapa parameter dipisah koma. Tipe bisa sama (a, b int) atau beda
//   (nama string, umur int). Urutan pemanggilan harus sesuai urutan parameter.
//
// Alasan penggunaan:
//   Fungsi butuh lebih dari satu input (nama+umur, panjang+lebar, dll).
//
// Fungsi penggunaan:
//   info("Siti", 22). Argumen pertama ke param pertama, kedua ke param kedua.
//
// Output:
//   Nama: Siti, Umur: 22
// -----------------------------------------------------------------------------
func banyakParameter() {
	info("Siti", 22)
}

func info(nama string, umur int) {
	fmt.Printf("  Nama: %s, Umur: %d\n", nama, umur)
}

// -----------------------------------------------------------------------------
// 3. Parameter tipe sama bisa disingkat
//
// Penjelasan:
//   func add(a int, b int) bisa ditulis func add(a, b int). Hanya tipe terakhir
//   yang ditulis; semua param sebelum itu dapat tipe yang sama.
//
// Alasan penggunaan:
//   Kode lebih ringkas tanpa mengubah arti. Umum untuk banyak param sejenis.
//
// Fungsi penggunaan:
//   add(3, 5). a=3, b=5, keduanya int.
//
// Output:
//   add(3, 5) = 8
// -----------------------------------------------------------------------------
func parameterTipeSama() {
	fmt.Printf("  add(3, 5) = %d\n", add(3, 5))
}

func add(a, b int) int {
	return a + b
}

// -----------------------------------------------------------------------------
// 4. Pass by value — nilai disalin ke parameter
//
// Penjelasan:
//   Di Go, argument ke fungsi secara default "pass by value": nilai disalin.
//   Perubahan di dalam fungsi pada parameter tidak mengubah variabel asli
//   di pemanggil (untuk int, string, struct value, dll).
//
// Alasan penggunaan:
//   Menghindari efek samping tak terduga: fungsi tidak mengubah data pemanggil
//   kecuali lewat return atau parameter pointer.
//
// Fungsi penggunaan:
//   x := 10; ubahCopy(x); // x tetap 10. Di dalam ubahCopy, param berubah saja.
//
// Output:
//   Setelah ubahCopy: x tetap 10
// -----------------------------------------------------------------------------
func passByValue() {
	x := 10
	ubahCopy(x)
	fmt.Printf("  Setelah ubahCopy: x tetap %d\n", x)
}

func ubahCopy(n int) {
	n = 99
}

// -----------------------------------------------------------------------------
// 5. Parameter pointer (*T) — mengubah nilai asli
//
// Penjelasan:
//   Parameter bertipe *T menerima alamat (pointer). Di dalam fungsi, *p mengakses
//   nilai yang ditunjuk. Mengubah *p mengubah variabel asli di pemanggil.
//
// Alasan penggunaan:
//   Saat fungsi harus mengubah nilai di tempat pemanggil (mutate), atau saat
//   struct besar dan ingin menghindari copy. Juga untuk optional (nil).
//
// Fungsi penggunaan:
//   x := 10; ubahPointer(&x); // x jadi 99. Pemanggil kirim &x (alamat).
//
// Output:
//   Sebelum: x = 10
//   Setelah ubahPointer: x = 99
// -----------------------------------------------------------------------------
func parameterPointer() {
	x := 10
	fmt.Printf("  Sebelum: x = %d\n", x)
	ubahPointer(&x)
	fmt.Printf("  Setelah ubahPointer: x = %d\n", x)
}

func ubahPointer(p *int) {
	*p = 99
}

// -----------------------------------------------------------------------------
// 6. Variadic parameter (...T)
//
// Penjelasan:
//   Parameter ...T menerima nol atau lebih argument bertipe T. Di dalam fungsi
//   param berperilaku seperti slice. Hanya satu parameter variadic per fungsi
//   dan harus di posisi terakhir.
//
// Alasan penggunaan:
//   Jumlah argument tidak tetap (sum(1,2,3) atau sum(1,2,3,4,5)) tanpa harus
//   membuat slice terlebih dahulu.
//
// Fungsi penggunaan:
//   total(1, 2, 3) atau total(slice...). Di dalam: for _, v := range nums.
//
// Output:
//   total(10, 20, 30) = 60
// -----------------------------------------------------------------------------
func parameterVariadic() {
	fmt.Printf("  total(10, 20, 30) = %d\n", total(10, 20, 30))
}

func total(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

// -----------------------------------------------------------------------------
// 7. Studi kasus: hitung total belanja (banyak parameter)
//
// Penjelasan:
//   Fungsi menerima beberapa harga item (parameter terpisah atau slice).
//   Di sini pakai variadic: totalBelanja(harga1, harga2, ...) return total.
//
// Alasan penggunaan:
//   Jumlah item belanja bisa berubah; variadic memudahkan pemanggilan tanpa
//   slice eksplisit.
//
// Fungsi penggunaan:
//   totalBelanja(5000, 10000, 7500) mengembalikan jumlah ketiganya.
//
// Output:
//   Total belanja (5000, 10000, 7500) = 22500
// -----------------------------------------------------------------------------
func kasusTotalBelanja() {
	fmt.Printf("  Total belanja (5000, 10000, 7500) = %d\n", totalBelanja(5000, 10000, 7500))
}

func totalBelanja(hargas ...int) int {
	sum := 0
	for _, h := range hargas {
		sum += h
	}
	return sum
}

// -----------------------------------------------------------------------------
// 8. Studi kasus: tukar nilai dua variabel dengan pointer
//
// Penjelasan:
//   Fungsi swap(a, b *int) menerima pointer ke dua int. Di dalam fungsi nilai
//   yang ditunjuk ditukar, sehingga variabel asli di pemanggil ikut tertukar.
//
// Alasan penggunaan:
//   Go tidak punya "pass by reference" untuk tipe dasar; pointer dipakai saat
//   fungsi harus mengubah lebih dari satu nilai "di tempat".
//
// Fungsi penggunaan:
//   x, y := 1, 2; swap(&x, &y); // x=2, y=1.
//
// Output:
//   Sebelum swap: a=1, b=2
//   Setelah swap: a=2, b=1
// -----------------------------------------------------------------------------
func kasusTukarNilai() {
	a, b := 1, 2
	fmt.Printf("  Sebelum swap: a=%d, b=%d\n", a, b)
	swap(&a, &b)
	fmt.Printf("  Setelah swap: a=%d, b=%d\n", a, b)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

// -----------------------------------------------------------------------------
// 9. Studi kasus: rata-rata banyak nilai (variadic)
//
// Penjelasan:
//   Fungsi rataRata(nums ...float64) mengembalikan rata-rata dari semua
//   argument. Jika tidak ada argument, return 0 (atau bisa return error).
//
// Alasan penggunaan:
//   Menghitung rata-rata dari N nilai tanpa harus bikin slice dulu; variadic
//   membuat pemanggilan natural.
//
// Fungsi penggunaan:
//   rataRata(70, 80, 90) return 80.0.
//
// Output:
//   rataRata(70, 80, 90) = 80.00
// -----------------------------------------------------------------------------
func kasusRataRataVariadic() {
	fmt.Printf("  rataRata(70, 80, 90) = %.2f\n", rataRata(70, 80, 90))
}

func rataRata(nums ...float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	sum := 0.0
	for _, n := range nums {
		sum += n
	}
	return sum / float64(len(nums))
}
