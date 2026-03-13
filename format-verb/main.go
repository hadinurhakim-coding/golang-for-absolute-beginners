// =============================================================================
// FORMAT VERB DI GO — Panduan fmt.Printf
// =============================================================================
// Format verb (format specifier) dipakai di:
//   - fmt.Printf(format, args...)  → cetak ke standar output
//   - fmt.Sprintf(format, args...) → kembalikan string (tidak cetak)
//   - fmt.Fprintf(w, format, args...) → tulis ke io.Writer (mis. file)
//
// Bentuk umum: %[flags][width][.precision]verb
//   flags    : opsional, mis. - untuk rata kiri, 0 untuk padding nol
//   width    : lebar minimal (jumlah karakter)
//   .precision: untuk float/string (jumlah digit atau karakter)
//   verb     : huruf yang menentukan tipe tampilan (d, f, s, v, dll)
// =============================================================================

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("========== 1. INTEGER: %d, %b, %o, %x, %X ==========")
	integerVerbs()

	fmt.Println("\n========== 2. FLOAT: %f, %e, %E, %g, %G ==========")
	floatVerbs()

	fmt.Println("\n========== 3. STRING & BOOLEAN: %s, %q, %t ==========")
	stringAndBoolVerbs()

	fmt.Println("\n========== 4. SERBA GUNA: %v, %+v, %#v, %T ==========")
	valueAndTypeVerbs()

	fmt.Println("\n========== 5. WIDTH & PRECISION (%5d, %.2f, %10s) ==========")
	widthAndPrecision()

	fmt.Println("\n========== 6. FLAGS: rata kiri (-), padding nol (0), tanda (+) ==========")
	flags()

	fmt.Println("\n========== 7. Sprintf & literal %% ==========")
	sprintfAndPercent()
}

// -----------------------------------------------------------------------------
// 1. Integer: %d (desimal), %b (biner), %o (oktal), %x/%X (heksadesimal)
// -----------------------------------------------------------------------------
func integerVerbs() {
	n := 42

	fmt.Printf("%%d (desimal)     : %d\n", n)   // 42
	fmt.Printf("%%b (biner)       : %b\n", n)   // 101010
	fmt.Printf("%%o (oktal)       : %o\n", n)   // 52
	fmt.Printf("%%x (hex kecil)   : %x\n", n)   // 2a
	fmt.Printf("%%X (hex besar)   : %X\n", n)   // 2A

	negatif := -42
	fmt.Printf("%%d untuk negatif : %d\n", negatif)
}

// -----------------------------------------------------------------------------
// 2. Float: %f (desimal), %e/%E (eksponen), %g/%G (pilih terpendek)
// -----------------------------------------------------------------------------
func floatVerbs() {
	pi := math.Pi

	fmt.Printf("%%f (desimal)     : %f\n", pi)       // default 6 digit di belakang koma
	fmt.Printf("%%.2f (2 digit)   : %.2f\n", pi)     // 3.14
	fmt.Printf("%%e (eksponen)    : %e\n", pi)       // 3.141593e+00
	fmt.Printf("%%E (E besar)     : %E\n", pi)       // 3.141593E+00
	fmt.Printf("%%g (singkat)     : %g\n", pi)       // Go pilih f atau e yang lebih ringkas
	fmt.Printf("%%g untuk besar   : %g\n", 1.23e10) // 1.23e+10
}

// -----------------------------------------------------------------------------
// 3. String (%s) dan Boolean (%t); %q untuk string dengan tanda kutip
// -----------------------------------------------------------------------------
func stringAndBoolVerbs() {
	s := "Hello, Go!"
	b := true

	fmt.Printf("%%s (string)      : %s\n", s)
	fmt.Printf("%%q (string+quote) : %q\n", s)   // "Hello, Go!" — berguna untuk debug
	fmt.Printf("%%t (boolean)     : %t\n", b)
	fmt.Printf("%%t false         : %t\n", false)
}

// -----------------------------------------------------------------------------
// 4. %v (value default), %+v (struct: tampilkan nama field), %#v (syntax Go), %T (tipe)
// -----------------------------------------------------------------------------
func valueAndTypeVerbs() {
	n := 42
	f := 3.14
	s := "hello"

	fmt.Printf("%%v (value)       : n=%v f=%v s=%v\n", n, f, s)
	fmt.Printf("%%T (type)        : n bertipe %T, f bertipe %T, s bertipe %T\n", n, f, s)

	// %#v: tampilan seperti literal Go (berguna untuk debug)
	fmt.Printf("%%#v (Go syntax)   : n=%#v s=%#v\n", n, s)

	type Orang struct {
		Nama string
		Umur int
	}
	o := Orang{Nama: "Budi", Umur: 20}
	fmt.Printf("%%v (struct)      : %v\n", o)
	fmt.Printf("%%+v (dengan field): %+v\n", o) // Nama:"Budi" Umur:20
	fmt.Printf("%%#v (struct)     : %#v\n", o)
}

// -----------------------------------------------------------------------------
// 5. Width (lebar minimal) dan precision (presisi / panjang)
// -----------------------------------------------------------------------------
func widthAndPrecision() {
	// Width: angka sebelum verb = lebar minimal (padding spasi di kiri, rata kanan)
	fmt.Printf("Tanpa width  : |%d|\n", 42)
	fmt.Printf("Width 5      : |%5d|\n", 42)   // "   42"
	fmt.Printf("Width 10     : |%10d|\n", 42)  // "        42"

	// Precision untuk float: angka setelah titik = digit di belakang koma
	fmt.Printf("%%.0f : %.0f\n", math.Pi)  // 3
	fmt.Printf("%%.2f : %.2f\n", math.Pi)  // 3.14
	fmt.Printf("%%.4f : %.4f\n", math.Pi)  // 3.1416

	// Width + precision untuk float: %width.precisionf
	fmt.Printf("|%8.2f|\n", math.Pi)  // "    3.14"

	// Width untuk string: batas panjang tampilan (atau padding)
	fmt.Printf("|%s|\n", "Hi")
	fmt.Printf("|%6s|\n", "Hi")   // "    Hi" (rata kanan)
	fmt.Printf("|%.2s|\n", "Hello") // "He" (potong 2 karakter)
}

// -----------------------------------------------------------------------------
// 6. Flags: - (rata kiri), 0 (padding nol), + (tampilkan tanda + untuk positif)
// -----------------------------------------------------------------------------
func flags() {
	fmt.Printf("Rata kanan (default) : |%5d|\n", 42)
	fmt.Printf("Rata kiri (-)         : |%-5d| selesai\n", 42)  // "42   " lalu " selesai"
	fmt.Printf("Padding nol (0)       : |%05d|\n", 42)          // 00042
	fmt.Printf("Tanda plus (+)        : %+d dan %+d\n", 42, -42) // +42 dan -42
}

// -----------------------------------------------------------------------------
// 7. Sprintf (hasil string, tidak cetak) dan cara cetak karakter %
// -----------------------------------------------------------------------------
func sprintfAndPercent() {
	// Sprintf mengembalikan string, tidak langsung cetak
	msg := fmt.Sprintf("Nilai: %d, Pi: %.2f", 100, math.Pi)
	fmt.Println("Hasil Sprintf:", msg)

	// Untuk mencetak karakter % literal, pakai %%
	fmt.Printf("Diskon 50%% hari ini.\n")        // Diskon 50% hari ini.
	fmt.Printf("Progress: %d%%\n", 75)           // Progress: 75%
}
