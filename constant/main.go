// =============================================================================
// KONSTANTA (CONSTANT) DI GO
// =============================================================================
// Konstanta: nilai yang tidak berubah setelah dideklarasikan.
// Harus bisa dihitung saat compile time (bukan hasil pemanggilan fungsi runtime).
// Deklarasi: const nama = nilai  atau  const nama tipe = nilai
// =============================================================================

package main // Paket utama agar file ini bisa dijalankan (go run)

import "fmt" // Untuk cetak ke layar (Println, Printf)

// main adalah titik masuk program. Isinya: cetak judul tiap bagian, lalu panggil fungsi yang mengurus materi.
func main() {
	fmt.Println("========== 1. Dasar: const tunggal & kelompok ==========") // Cetak judul bagian 1
	dasar()                                                                 // Jalankan fungsi dasar()

	fmt.Println("\n========== 2. Konstanta bertipe vs tak bertipe ==========") // \n = pindah baris dulu
	typedVsUntyped()

	fmt.Println("\n========== 3. iota — konstanta berurutan ==========")
	iotaDasar()

	fmt.Println("\n========== 4. iota lanjutan: shift & skip ==========")
	iotaLanjutan()

	fmt.Println("\n========== 5. Konstanta untuk ukuran array ==========")
	constUntukArray()

	fmt.Println("\n========== 6. Konstanta vs variabel ==========")
	konstantaVsVariabel()
}

// -----------------------------------------------------------------------------
// 1. const nama = nilai (inferensi tipe) atau const nama tipe = nilai
//    Boleh dikelompokkan dalam ( )
// -----------------------------------------------------------------------------
func dasar() {
	// Konstanta tanpa tipe eksplisit: Go tentukan tipe dari nilai (3.14159 → seperti float64)
	const Pi = 3.14159
	// Konstanta dengan tipe eksplisit: AppName selalu string
	const AppName string = "MyApp"

	// %v = tampilkan nilai apa adanya; %q = tampilkan dengan tanda kutip (untuk string)
	fmt.Printf("Pi = %v (untyped, dipakai sebagai float64)\n", Pi)
	fmt.Printf("AppName = %q\n", AppName)

	// Beberapa konstanta bisa dideklarasikan dalam satu blok ( )
	const (
		MaxUsers = 100 // Nilai 100
		MinAge   = 18  // Nilai 18 (spasi sebelum 18 hanya untuk rapi)
	)
	fmt.Printf("MaxUsers = %d, MinAge = %d\n", MaxUsers, MinAge) // %d = bilangan bulat
}

// -----------------------------------------------------------------------------
// 2. Typed: const x int = 42 → tipe tetap int (konversi ke tipe lain harus eksplisit)
//    Untyped: const x = 42 → tipe tergantung konteks (bisa int, float64, dll)
// -----------------------------------------------------------------------------
func typedVsUntyped() {
	const typed int = 42   // Konstanta BERTIPE: tipe tetap int, tidak bisa dipakai sebagai float tanpa konversi
	const untyped = 42     // Konstanta TAK BERTIPE: tipe mengikuti konteks (bisa int, float64, dll)

	var a int = typed               // Masuk ke int → OK
	var b float64 = float64(typed)  // Ke float64 harus konversi eksplisit: float64(typed)
	var c float64 = untyped         // untyped 42 boleh dipakai sebagai float64 tanpa konversi

	fmt.Printf("typed=%d (tipe int), untyped dipakai sebagai float64: %v\n", typed, c)
	_ = a // Pakai variabel a agar compiler tidak komplain "declared and not used"
	_ = b
}

// -----------------------------------------------------------------------------
// 3. iota: nilai integer yang naik otomatis per baris (mulai 0)
//    Reset di setiap blok const ( )
// -----------------------------------------------------------------------------
func iotaDasar() {
	// iota = bilangan integer yang naik otomatis: 0, 1, 2, ... di dalam satu blok const
	const (
		A = iota // Baris pertama: iota = 0, jadi A = 0
		B        // Baris kedua: iota = 1. Karena tidak tulis "= iota", pakai aturan yang sama → B = 1
		C        // Baris ketiga: iota = 2 → C = 2
	)
	fmt.Printf("A=%d, B=%d, C=%d\n", A, B, C)

	// Setiap blok const ( ) baru, iota di-reset mulai dari 0 lagi
	const (
		Senin  = iota // Baris pertama blok ini: iota = 0 → Senin = 0
		Selasa        // iota = 1 → Selasa = 1
		Rabu          // iota = 2 → Rabu = 2
	)
	fmt.Printf("Senin=%d, Selasa=%d, Rabu=%d\n", Senin, Selasa, Rabu)
}

// -----------------------------------------------------------------------------
// 4. iota dengan ekspresi (shift, skip dengan _)
// -----------------------------------------------------------------------------
func iotaLanjutan() {
	// iota bisa dipakai dalam ekspresi. << = bit shift ke kiri (pangkat 2)
	// 1 << 0 = 1, 1 << 10 = 1024, 1 << 20 = 1048576 (sering untuk KB, MB, GB)
	const (
		KB = 1 << (10 * iota) // Baris 0: 10*iota=0 → 1<<0 = 1
		MB                     // Baris 1: 10*iota=10 → 1<<10 = 1024
		GB                     // Baris 2: 10*iota=20 → 1<<20 = 1048576
	)
	fmt.Printf("KB=%d, MB=%d, GB=%d\n", KB, MB, GB)

	// _ (blank identifier) = kita tidak pakai nilai itu; iota tetap naik
	const (
		Low    = iota // 0
		_             // iota=1, tapi tidak disimpan ke nama (skip)
		Medium        // iota=2 → Medium = 2
		_             // iota=3, skip
		High          // iota=4 → High = 4
	)
	fmt.Printf("Low=%d, Medium=%d, High=%d\n", Low, Medium, High)
}

// -----------------------------------------------------------------------------
// 5. Konstanta bisa dipakai untuk ukuran array (variabel tidak boleh)
// -----------------------------------------------------------------------------
func constUntukArray() {
	const N = 3                              // Ukuran array harus known at compile time
	var arr [N]int = [N]int{1, 2, 3}         // Array 3 elemen int: arr[0]=1, arr[1]=2, arr[2]=3
	fmt.Printf("Array dengan size const N=%d: %v\n", N, arr) // %v cetak slice/array
}

// -----------------------------------------------------------------------------
// 6. Ringkasan: konstanta harus known at compile time
// -----------------------------------------------------------------------------
func konstantaVsVariabel() {
	const C = 10 // Nilai 10 diketahui saat compile → boleh
	// const D = getNumber()  // Error: getNumber() jalan saat runtime, bukan compile time

	fmt.Println("Konstanta harus bisa dihitung saat compile.")
	fmt.Println("Variabel bisa dari input/rumus runtime.")
	fmt.Printf("Contoh const C = %d\n", C)
}
