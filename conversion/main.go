// =============================================================================
// KONVERSI TIPE DATA (TYPE CONVERSION) DI GO
// =============================================================================
// Di Go, konversi tipe harus EKSPLISIT: T(v) = ubah nilai v ke tipe T.
// Tidak ada konversi implisit (otomatis) antar tipe yang berbeda — termasuk number.
// Contoh: int(3.14), float64(42), string([]byte{...})
// =============================================================================

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("========== 1. Konversi number: int ↔ float ==========")
	numberConversion()

	fmt.Println("\n========== 2. Konversi number: int ↔ int (ukuran beda) ==========")
	intSizes()

	fmt.Println("\n========== 3. Float ke int: truncate (bukan pembulatan) ==========")
	floatToInt()

	fmt.Println("\n========== 4. Number ke string & string ke number (strconv) ==========")
	strconvConversion()

	fmt.Println("\n========== 5. string ↔ []byte (slice byte) ==========")
	stringByteConversion()

	fmt.Println("\n========== 6. Konversi yang invalid / bahaya ==========")
	invalidConversion()
}

// -----------------------------------------------------------------------------
// 1. int ke float64, float64 ke int — harus tulis tipe tujuan secara eksplisit
// -----------------------------------------------------------------------------
func numberConversion() {
	i := 42
	f := 3.14

	// int → float64: tambah presisi desimal (aman)
	fFromInt := float64(i)
	fmt.Printf("int %d → float64 = %v\n", i, fFromInt)

	// float64 → int: nilai di belakang koma hilang (truncate)
	iFromFloat := int(f)
	fmt.Printf("float64 %v → int = %d (bukan 3.14 dibulatkan, tapi 3)\n", f, iFromFloat)
}

// -----------------------------------------------------------------------------
// 2. int ↔ int8, int16, int32, int64 — konversi ukuran berbeda
//    Nilai di luar rentang tipe tujuan = overflow (nilai berubah)
// -----------------------------------------------------------------------------
func intSizes() {
	var i32 int32 = 100
	var i64 int64 = 1000

	fmt.Printf("int32 %d → int64 = %d\n", i32, int64(i32))
	fmt.Printf("int64 %d → int32 = %d\n", i64, int32(i64))

	// Overflow: nilai terlalu besar untuk int8 (-128..127)
	big := 300
	fmt.Printf("300 → int8 = %d (overflow, nilai 'melipat')\n", int8(big))
}

// -----------------------------------------------------------------------------
// 3. float ke int: selalu truncate (potong), bukan round (pembulatan)
//    Untuk pembulatan pakai math.Round() dulu, baru int()
// -----------------------------------------------------------------------------
func floatToInt() {
	f1, f2 := 3.7, 3.2
	fmt.Printf("int(3.7) = %d (bukan 4)\n", int(f1))
	fmt.Printf("int(3.2) = %d\n", int(f2))
}

// -----------------------------------------------------------------------------
// 4. strconv: number ↔ string
//    Itoa = int to string; Atoi = string to int (kembalian int, error)
//    FormatFloat, ParseFloat untuk float; FormatInt, ParseInt untuk kontrol basis
// -----------------------------------------------------------------------------
func strconvConversion() {
	// int → string
	n := 42
	s := strconv.Itoa(n)
	fmt.Printf("strconv.Itoa(42) = %q\n", s)

	// string → int (Atoi mengembalikan int dan error)
	s2 := "100"
	n2, err := strconv.Atoi(s2)
	if err != nil {
		fmt.Println("Atoi error:", err)
	} else {
		fmt.Printf("strconv.Atoi(%q) = %d\n", s2, n2)
	}

	// float → string (format: 'f'=desimal, prec=jumlah digit)
	f := 3.14159
	sf := strconv.FormatFloat(f, 'f', 2, 64) // 2 digit di belakang koma, float64
	fmt.Printf("FormatFloat(3.14159, 'f', 2, 64) = %q\n", sf)

	// string → float64
	parsed, _ := strconv.ParseFloat("2.5", 64)
	fmt.Printf("ParseFloat(\"2.5\", 64) = %v\n", parsed)
}

// -----------------------------------------------------------------------------
// 5. string dan []byte bisa dikonversi langsung: []byte(s), string(bytes)
//    Berguna untuk baca/tulis data biner atau modifikasi "string" lewat slice byte
// -----------------------------------------------------------------------------
func stringByteConversion() {
	s := "Go"
	bytes := []byte(s)
	fmt.Printf("[]byte(%q) = %v\n", s, bytes)

	back := string(bytes)
	fmt.Printf("string([]byte) = %q\n", back)
}

// -----------------------------------------------------------------------------
// 6. Yang tidak boleh / hati-hati
//    - bool tidak dikonversi dari int (0/1) seperti di C
//    - string dan number tidak dikonversi langsung; pakai strconv
// -----------------------------------------------------------------------------
func invalidConversion() {
	// var b bool = bool(1)  // Error: cannot convert 1 to bool
	fmt.Println("bool tidak bisa dikonversi dari int (tidak ada bool(1) di Go).")
	fmt.Println("string dan int tidak bisa konversi langsung: pakai strconv.Itoa / Atoi.")
}
