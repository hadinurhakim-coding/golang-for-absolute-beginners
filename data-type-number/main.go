// =============================================================================
// TIPE DATA NUMBER DI GO — Panduan Lengkap
// =============================================================================
// Di Go, tipe data number terbagi menjadi:
//   1. Integer (bilangan bulat) — signed (bertanda) dan unsigned (tak bertanda)
//   2. Floating-point (bilangan desimal)
//   3. Complex (bilangan kompleks)
// Go TIDAK punya tipe "float" atau "double" seperti di bahasa lain; nama tipe eksplisit.
// =============================================================================
//
// =============================================================================
// STRUKTUR KODE — Mengapa Ada Banyak func?
// =============================================================================
//
// 1. ALUR PROGRAM
//    - Go hanya punya SATU titik masuk: func main(). Saat program jalan, yang
//      pertama dipanggil adalah main().
//    - Di dalam main() kita TIDAK menulis semua kode sekaligus. Kita memanggil
//      fungsi-fungsi lain (integerTypes(), floatingPointTypes(), dll) supaya
//      main() tetap singkat dan mudah dibaca.
//
// 2. MENGAPA BANYAK FUNGSI?
//    - Satu fungsi = satu tanggung jawab (satu topik). Ini namanya pemisahan
//      tanggung jawab (separation of concerns). integerTypes() hanya urusan
//      integer; floatingPointTypes() hanya float; dan seterusnya.
//    - Lebih mudah dicari: mau ubah contoh float? Buka floatingPointTypes().
//    - Lebih mudah diuji: nanti Anda bisa mengetes tiap fungsi sendiri.
//    - main() jadi seperti "daftar isi": hanya memanggil fungsi lain sesuai
//      urutan yang kita mau.
//
// 3. BENTUK UMUM
//
//    func main() {
//        // 1. Cetak judul bagian
//        // 2. Panggil fungsi yang mengurus bagian itu
//        integerTypes()      // <- fungsi tanpa parameter, tanpa return
//        floatingPointTypes()
//        ...
//    }
//
//    func integerTypes() {
//        // Variabel dan logika HANYA untuk integer; tidak campur dengan float.
//    }
//
// 4. URUTAN EKSEKUSI (saat Anda go run)
//    main() jalan
//      → fmt.Println("========== 1. ...")  (cetak judul)
//      → integerTypes()                     (jalan sampai selesai)
//      → fmt.Println("\n========== 2. ...")
//      → integerSizesAndRanges()
//      → ... dan seterusnya sampai zeroValues() selesai
//    → main() selesai → program selesai
//
// 5. FUNGSI DI FILE INI TIDAK PUNYA PARAMETER ATAU RETURN
//    Semua func di sini bentuknya: func namaFungsi() { ... }
//    Artinya: tidak menerima input, tidak mengembalikan nilai. Hanya menjalankan
//    kode di dalamnya (variabel + fmt.Printf). Ini cukup untuk tujuan belajar.
//    Nanti Anda akan belajar func dengan parameter dan return value.
//
// =============================================================================

package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	fmt.Println("========== 1. INTEGER (Bilangan Bulat) ==========")
	integerTypes()

	fmt.Println("\n========== 2. INTEGER: Ukuran & Rentang ==========")
	integerSizesAndRanges()

	fmt.Println("\n========== 3. int vs int32/int64 ==========")
	intVsFixedSize()

	fmt.Println("\n========== 4. FLOATING-POINT (Desimal) ==========")
	floatingPointTypes()

	fmt.Println("\n========== 5. COMPLEX (Bilangan Kompleks) ==========")
	complexTypes()

	fmt.Println("\n========== 6. KONVERSI TIPE & LITERAL ==========")
	conversionAndLiterals()

	fmt.Println("\n========== 7. ZERO VALUE ==========")
	zeroValues()
}

// -----------------------------------------------------------------------------
// 1. INTEGER: signed (bisa negatif) dan unsigned (hanya >= 0)
// -----------------------------------------------------------------------------
func integerTypes() {
	// Signed: bisa negatif, nol, atau positif
	var a int8 = 127   // -128 sampai 127
	var b int16 = 32000
	var c int32 = 2_000_000_000   // underscore boleh untuk keterbacaan
	var d int64 = 9_000_000_000_000_000_000

	// Unsigned: hanya >= 0
	var e uint8 = 255  // 0 sampai 255 (sering dipakai sebagai "byte")
	var f uint16 = 65535
	var g uint32 = 4_000_000_000
	var h uint64 = 18_000_000_000_000_000_000

	fmt.Printf("Signed:   int8=%d  int16=%d  int32=%d  int64=%d\n", a, b, c, d)
	fmt.Printf("Unsigned: uint8=%d  uint16=%d  uint32=%d  uint64=%d\n", e, f, g, h)

	// uint8 alias: "byte" (sering untuk data biner/string)
	var by byte = 255
	fmt.Printf("byte (alias uint8) = %d\n", by)
}

// -----------------------------------------------------------------------------
// 2. Ukuran (byte) dan rentang nilai min–max
// -----------------------------------------------------------------------------
func integerSizesAndRanges() {
	// unsafe.Sizeof mengembalikan ukuran dalam byte
	fmt.Printf("int8   : %d byte, rentang %d ... %d\n", unsafe.Sizeof(int8(0)), math.MinInt8, math.MaxInt8)
	fmt.Printf("int16  : %d byte, rentang %d ... %d\n", unsafe.Sizeof(int16(0)), math.MinInt16, math.MaxInt16)
	fmt.Printf("int32  : %d byte, rentang %d ... %d\n", unsafe.Sizeof(int32(0)), math.MinInt32, math.MaxInt32)
	fmt.Printf("int64  : %d byte, rentang %d ... %d\n", unsafe.Sizeof(int64(0)), math.MinInt64, math.MaxInt64)

	fmt.Printf("uint8  : %d byte, rentang 0 ... %d\n", unsafe.Sizeof(uint8(0)), math.MaxUint8)
	fmt.Printf("uint16 : %d byte, rentang 0 ... %d\n", unsafe.Sizeof(uint16(0)), math.MaxUint16)
	fmt.Printf("uint32 : %d byte, rentang 0 ... %d\n", unsafe.Sizeof(uint32(0)), math.MaxUint32)
	fmt.Printf("uint64 : %d byte (rentang 0 ... 2^64-1, terlalu besar untuk dicetak di sini)\n", unsafe.Sizeof(uint64(0)))
}

// -----------------------------------------------------------------------------
// 3. int dan uint — ukuran mengikuti platform (32 atau 64 bit)
// -----------------------------------------------------------------------------
func intVsFixedSize() {
	var i int   // di Windows 64-bit = 8 byte, di ARM 32-bit = 4 byte
	var u uint

	fmt.Printf("int  : %d byte (platform-dependent)\n", unsafe.Sizeof(i))
	fmt.Printf("uint : %d byte (platform-dependent)\n", unsafe.Sizeof(u))

	// Kapan pakai int? Untuk indeks, panjang slice, loop — biasanya cukup "int".
	// Kapan pakai int32/int64? Kalau butuh ukuran pasti (serialisasi, protokol, API).
}

// -----------------------------------------------------------------------------
// 4. FLOATING-POINT — float32 dan float64
// -----------------------------------------------------------------------------
func floatingPointTypes() {
	var f32 float32 = 3.141592653589793
	var f64 float64 = 3.141592653589793

	fmt.Printf("float32: %d byte, contoh %.6f (presisi terbatas)\n", unsafe.Sizeof(f32), f32)
	fmt.Printf("float64: %d byte, contoh %.15f (presisi lebih baik)\n", unsafe.Sizeof(f64), f64)

	// Literal desimal tanpa suffix = float64
	x := 1.5         // tipe: float64
	var y float32 = 1.5

	fmt.Printf("Literal 1.5 default = float64 (%v). Eksplisit float32: %v\n", x, y)

	// Konstanta berguna
	fmt.Printf("math.Pi (float64) = %.15f\n", math.Pi)
	fmt.Printf("math.MaxFloat32 = %e\n", math.MaxFloat32)
	fmt.Printf("math.SmallestNonzeroFloat32 = %e\n", math.SmallestNonzeroFloat32)
}

// -----------------------------------------------------------------------------
// 5. COMPLEX — complex64 (float32+float32) dan complex128 (float64+float64)
// -----------------------------------------------------------------------------
func complexTypes() {
	// complex(real, imaginer)
	var c64 complex64 = complex(1, 2)   // 1 + 2i
	c128 := complex(1.0, 2.0)           // tipe: complex128 (default)

	fmt.Printf("complex64  : %d byte, nilai %v\n", unsafe.Sizeof(c64), c64)
	fmt.Printf("complex128 : %d byte, nilai %v\n", unsafe.Sizeof(c128), c128)

	// Akses bagian real dan imajiner
	fmt.Printf("real(c128)=%v, imag(c128)=%v\n", real(c128), imag(c128))
}

// -----------------------------------------------------------------------------
// 6. Konversi tipe (harus eksplisit) & literal number
// -----------------------------------------------------------------------------
func conversionAndLiterals() {
	// Go TIDAK mengkonversi otomatis antar tipe number
	var a int32 = 100
	// b := int64(a)  // benar: konversi eksplisit
	var b int64 = int64(a)
	fmt.Printf("int32 100 → int64 = %d\n", b)

	// Literal: bilangan bulat default int; desimal default float64
	n := 42        // int
	nf := 42.0     // float64
	fmt.Printf("42 → tipe int; 42.0 → float64. n=%d, nf=%v\n", n, nf)

	// Konversi float → int (nilai di belakang koma hilang)
	f := 3.99
	fmt.Printf("float64 3.99 → int = %d (bukan pembulatan, tapi truncate)\n", int(f))
}

// -----------------------------------------------------------------------------
// 7. Zero value: nilai default saat variabel belum diisi
// -----------------------------------------------------------------------------
func zeroValues() {
	var i8 int8
	var u8 uint8
	var f64 float64
	var c128 complex128

	fmt.Printf("Zero value int8      = %d\n", i8)
	fmt.Printf("Zero value uint8     = %d\n", u8)
	fmt.Printf("Zero value float64   = %v\n", f64)
	fmt.Printf("Zero value complex128= %v\n", c128)
}
