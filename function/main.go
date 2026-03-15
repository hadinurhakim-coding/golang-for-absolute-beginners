// =============================================================================
// FUNCTION (FUNGSI) DI GO
// =============================================================================
// Function = blok kode yang bisa dipanggil berulang. Mengurangi duplikasi dan
// memecah program jadi bagian kecil. Bentuk: func nama(param tipe) returnTipe { }
//
// POLA KOMENTAR (gunakan konsisten di direktori lain):
//   1. Penjelasan   — ringkasan materi / apa yang dilakukan.
//   2. Alasan penggunaan — kenapa fitur ini dipakai.
//   3. Fungsi penggunaan — cara pakai (pemanggilan, parameter, return).
//   4. Output       — contoh hasil cetakan/return dari fungsi tersebut.
// =============================================================================

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("========== 1. Fungsi tanpa parameter & tanpa return ==========")
	fungsiDasar()

	fmt.Println("\n========== 2. Fungsi dengan parameter ==========")
	fungsiParameter()

	fmt.Println("\n========== 3. Fungsi dengan return value ==========")
	fungsiReturn()

	fmt.Println("\n========== 4. Multiple return values ==========")
	fungsiMultipleReturn()

	fmt.Println("\n========== 5. Named return values ==========")
	fungsiNamedReturn()

	fmt.Println("\n========== 6. Variadic function (parameter ...T) ==========")
	fungsiVariadic()

	fmt.Println("\n========== 7. Studi kasus: luas persegi panjang ==========")
	kasusLuasPersegiPanjang()

	fmt.Println("\n========== 8. Studi kasus: cek bilangan prima ==========")
	kasusCekPrima()

	fmt.Println("\n========== 9. Studi kasus: format nama (multiple return) ==========")
	kasusFormatNama()

	fmt.Println("\n========== 10. Studi kasus: jumlah banyak angka (variadic) ==========")
	kasusJumlahVariadic()
}

// -----------------------------------------------------------------------------
// 1. Fungsi tanpa parameter dan tanpa return
//
// Penjelasan:
//   Fungsi paling sederhana: tidak menerima input, tidak mengembalikan nilai.
//   Hanya menjalankan blok kode saat dipanggil.
//
// Alasan penggunaan:
//   Untuk kode yang hanya perlu dijalankan (cetak, inisialisasi, side effect)
//   tanpa perlu mengirim/menerima data.
//
// Fungsi penggunaan:
//   Dipanggil dengan nama fungsi + (). Berguna untuk mengelompokkan kode
//   yang berulang atau memisahkan langkah dalam program.
//
// Output:
//   Hello dari fungsi dasar!
// -----------------------------------------------------------------------------
func fungsiDasar() {
	fmt.Println("  Hello dari fungsi dasar!")
}

// -----------------------------------------------------------------------------
// 2. Fungsi dengan parameter
//
// Penjelasan:
//   Parameter = input yang diterima fungsi. Tiap parameter punya nama dan tipe.
//   func nama(param1 T1, param2 T2) — bisa satu atau banyak parameter.
//
// Alasan penggunaan:
//   Agar fungsi bisa bekerja dengan data yang berbeda-beda setiap pemanggilan
//   tanpa mengubah kode di dalam fungsi.
//
// Fungsi penggunaan:
//   Memanggil dengan nilai konkret: greet("Budi"). Nilai disalin ke parameter
//   (pass by value untuk tipe dasar).
//
// Output:
//   Halo, Budi!
//   Umur: 20
// -----------------------------------------------------------------------------
func fungsiParameter() {
	greet("Budi")
	umur(20)
}

func greet(nama string) {
	fmt.Printf("  Halo, %s!\n", nama)
}

func umur(u int) {
	fmt.Printf("  Umur: %d\n", u)
}

// -----------------------------------------------------------------------------
// 3. Fungsi dengan return value
//
// Penjelasan:
//   Return value = nilai yang dikembalikan ke pemanggil. Tipe return ditulis
//   setelah parameter: func nama() T. Pakai return nilai untuk mengembalikan.
//
// Alasan penggunaan:
//   Agar hasil perhitungan atau proses bisa dipakai di tempat pemanggilan
//   (assign ke variabel, dipakai di ekspresi, dikirim ke fungsi lain).
//
// Fungsi penggunaan:
//   result := add(3, 5). Fungsi harus return nilai yang tipenya sesuai.
//
// Output:
//   add(3, 5) = 8
//   double(7) = 14
// -----------------------------------------------------------------------------
func fungsiReturn() {
	fmt.Printf("  add(3, 5) = %d\n", add(3, 5))
	fmt.Printf("  double(7) = %d\n", double(7))
}

func add(a, b int) int {
	return a + b
}

func double(x int) int {
	return x * 2
}

// -----------------------------------------------------------------------------
// 4. Multiple return values
//
// Penjelasan:
//   Go mengizinkan fungsi mengembalikan lebih dari satu nilai. Bentuk:
//   func nama() (T1, T2). Return: return v1, v2. Pemanggil: a, b := nama().
//
// Alasan penggunaan:
//   Sering dipakai untuk nilai + error (hasil, err). Juga untuk pasangan
//   nilai (hasil bagi, sisa; min, max; dll) tanpa perlu struct.
//
// Fungsi penggunaan:
//   quotient, remainder := divide(10, 3). Bisa ignore dengan _: v, _ := divide(...).
//
// Output:
//   divide(10, 3) = 3 sisa 1
//   split("halo-dunia", "-") = [halo dunia]
// -----------------------------------------------------------------------------
func fungsiMultipleReturn() {
	q, r := divide(10, 3)
	fmt.Printf("  divide(10, 3) = %d sisa %d\n", q, r)
	a, b := split("halo-dunia", "-")
	fmt.Printf("  split(\"halo-dunia\", \"-\") = [%s %s]\n", a, b)
}

func divide(a, b int) (int, int) {
	return a / b, a % b
}

func split(s, sep string) (string, string) {
	parts := strings.Split(s, sep)
	if len(parts) < 2 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}

// -----------------------------------------------------------------------------
// 5. Named return values
//
// Penjelasan:
//   Return value bisa diberi nama di signature: func nama() (x int, y int).
//   Variabel tersebut dianggap dideklarasikan di awal fungsi; return tanpa
//   nilai mengembalikan nilai terakhir variabel tersebut (naked return).
//
// Alasan penggunaan:
//   Dokumentasi di signature; kadang kode lebih singkat. Hati-hati di fungsi
//   panjang agar tidak membingungkan.
//
// Fungsi penggunaan:
//   return saja (naked return) akan mengembalikan nilai current dari nama tersebut.
//   Bisa tetap return eksplisit: return 1, 2.
//
// Output:
//   minMax(3, 7) = min 3, max 7
// -----------------------------------------------------------------------------
func fungsiNamedReturn() {
	mn, mx := minMax(3, 7)
	fmt.Printf("  minMax(3, 7) = min %d, max %d\n", mn, mx)
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
// 6. Variadic function (parameter ...T)
//
// Penjelasan:
//   Parameter variadic menerima nol atau lebih nilai dengan tipe yang sama.
//   Bentuk: func nama(nums ...int). Di dalam fungsi, nums berperilaku seperti slice.
//
// Alasan penggunaan:
//   Saat jumlah argumen tidak tetap (jumlah angka, banyak string, dll) tanpa
//   harus membuat slice lalu memanggil fungsi dengan slice tersebut.
//
// Fungsi penggunaan:
//   sum(1, 2, 3) atau sum(slice...) untuk melebarkan slice jadi argumen.
//
// Output:
//   sum(1, 2, 3, 4, 5) = 15
//   join(", ", "A", "B", "C") = A, B, C
// -----------------------------------------------------------------------------
func fungsiVariadic() {
	fmt.Printf("  sum(1, 2, 3, 4, 5) = %d\n", sum(1, 2, 3, 4, 5))
	fmt.Printf("  join(\", \", \"A\", \"B\", \"C\") = %s\n", join(", ", "A", "B", "C"))
}

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func join(sep string, parts ...string) string {
	return strings.Join(parts, sep)
}

// -----------------------------------------------------------------------------
// 7. Studi kasus: luas persegi panjang
//
// Penjelasan:
//   Fungsi menerima panjang dan lebar, mengembalikan luas (panjang * lebar).
//
// Alasan penggunaan:
//   Menghitung luas dipakai di banyak tempat; dengan fungsi cukup panggil
//   luasPersegiPanjang(p, l) dan dapat satu nilai return.
//
// Fungsi penggunaan:
//   luas := luasPersegiPanjang(5, 3) lalu cetak atau pakai untuk perhitungan lain.
//
// Output:
//   Luas persegi panjang 5 x 3 = 15
// -----------------------------------------------------------------------------
func kasusLuasPersegiPanjang() {
	p, l := 5.0, 3.0
	luas := luasPersegiPanjang(p, l)
	fmt.Printf("  Luas persegi panjang %.0f x %.0f = %.0f\n", p, l, luas)
}

func luasPersegiPanjang(panjang, lebar float64) float64 {
	return panjang * lebar
}

// -----------------------------------------------------------------------------
// 8. Studi kasus: cek bilangan prima
//
// Penjelasan:
//   Fungsi menerima satu integer, mengembalikan true jika prima, false jika tidak.
//   Prima = hanya habis dibagi 1 dan dirinya sendiri (cek pembagi 2 sampai sqrt(n)).
//
// Alasan penggunaan:
//   Logika cek prima dipakai berulang; dengan fungsi bisa dipanggil dari mana saja
//   dan return bool untuk dipakai di if atau variabel.
//
// Fungsi penggunaan:
//   if cekPrima(7) { ... }. Return value dipakai langsung di kondisi atau assign.
//
// Output:
//   7 prima? true
//   10 prima? false
// -----------------------------------------------------------------------------
func kasusCekPrima() {
	fmt.Printf("  7 prima? %t\n", cekPrima(7))
	fmt.Printf("  10 prima? %t\n", cekPrima(10))
}

func cekPrima(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// -----------------------------------------------------------------------------
// 9. Studi kasus: format nama (depan, belakang) — multiple return
//
// Penjelasan:
//   Fungsi menerima nama lengkap (string), mengembalikan nama depan dan nama
//   belakang. Asumsi dipisah spasi; jika satu kata, belakang kosong.
//
// Alasan penggunaan:
//   Satu string perlu dipecah jadi dua nilai untuk form, tampilan, atau API;
//   multiple return menghindari struct sederhana atau parsing berulang.
//
// Fungsi penggunaan:
//   depan, belakang := formatNama("Budi Santoso"). Bisa _ untuk ignore salah satu.
//
// Output:
//   formatNama("Budi Santoso") = depan: Budi, belakang: Santoso
// -----------------------------------------------------------------------------
func kasusFormatNama() {
	depan, belakang := formatNama("Budi Santoso")
	fmt.Printf("  formatNama(\"Budi Santoso\") = depan: %s, belakang: %s\n", depan, belakang)
}

func formatNama(namaLengkap string) (depan, belakang string) {
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
// 10. Studi kasus: jumlah banyak angka (variadic)
//
// Penjelasan:
//   Fungsi variadic yang menjumlahkan semua angka yang dikirim. Contoh
//   sum(1,2,3) dan sum(10,20,30,40). Menggunakan parameter ...int.
//
// Alasan penggunaan:
//   Jumlah angka bisa berubah-ubah (3 nilai, 5 nilai, dll); variadic membuat
//   pemanggilan natural tanpa harus membuat slice dulu.
//
// Fungsi penggunaan:
//   total := sum(10, 20, 30, 40). Atau slice: total := sum(nums...).
//
// Output:
//   sum(10, 20, 30, 40) = 100
// -----------------------------------------------------------------------------
func kasusJumlahVariadic() {
	fmt.Printf("  sum(10, 20, 30, 40) = %d\n", sum(10, 20, 30, 40))
}
