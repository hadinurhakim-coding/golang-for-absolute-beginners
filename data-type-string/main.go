// =============================================================================
// TIPE DATA STRING DI GO
// =============================================================================
// String di Go:
//   - Tipe: string (bukan []byte, tapi bisa dikonversi)
//   - Immutable: isi tidak bisa diubah per karakter; buat string baru untuk "ubah"
//   - Encoding: UTF-8 by default; satu karakter bisa 1–4 byte (rune = 1 Unicode code point)
//   - Literal: "double quote" (boleh \n \t \"), `backtick` (raw, tidak ada escape)
// =============================================================================

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println("========== 1. Dasar: deklarasi & literal ==========")
	dasar()

	fmt.Println("\n========== 2. Raw string (backtick) & escape ==========")
	rawDanEscape()

	fmt.Println("\n========== 3. Zero value & string kosong ==========")
	zeroValue()

	fmt.Println("\n========== 4. Penggabungan string (concatenation) ==========")
	penggabungan()

	fmt.Println("\n========== 5. Panjang: len() vs jumlah karakter (rune) ==========")
	panjangString()

	fmt.Println("\n========== 6. Akses per karakter: byte, rune, range ==========")
	aksesKarakter()

	fmt.Println("\n========== 7. Fungsi dari paket strings ==========")
	paketStrings()

	fmt.Println("\n========== 8. String immutable ==========")
	immutable()

	fmt.Println("\n========== 9. fmt.Sprintf untuk format string ==========")
	formatString()
}

// -----------------------------------------------------------------------------
// 1. Deklarasi: var s string = "...", s := "..."
// -----------------------------------------------------------------------------
func dasar() {
	var s1 string = "Hello, Go!"
	s2 := "Inferensi tipe"

	fmt.Printf("s1 = %q\n", s1)
	fmt.Printf("s2 = %q\n", s2)
	fmt.Printf("Tipe s1: %T\n", s1)
}

// -----------------------------------------------------------------------------
// 2. Raw string dengan backtick `...` — tidak ada escape; newline tetap
//    Escape di "..." : \n (newline), \t (tab), \" (kutip), \\ (backslash)
// -----------------------------------------------------------------------------
func rawDanEscape() {
	// Double quote: pakai escape untuk newline dan kutip
	escaped := "Baris 1\nBaris 2\t(tab)\nKutip: \"ini dalam kutip\""
	fmt.Println(escaped)

	// Backtick: raw, semua karakter literal (termasuk newline)
	raw := `Ini raw string.
Newline tetap.
Backslash \ dan "kutip" tidak di-escape.`
	fmt.Println(raw)
}

// -----------------------------------------------------------------------------
// 3. Zero value string = "" (string kosong)
// -----------------------------------------------------------------------------
func zeroValue() {
	var s string
	fmt.Printf("Zero value string = %q (panjang %d)\n", s, len(s))
	fmt.Printf("s == \"\" → %t\n", s == "")
}

// -----------------------------------------------------------------------------
// 4. Penggabungan: + dan strings.Builder (untuk banyak penggabungan)
// -----------------------------------------------------------------------------
func penggabungan() {
	a, b := "Hello", " Go!"
	fmt.Printf("a + b = %q\n", a+b)

	// Banyak penggabungan: strings.Join atau strings.Builder
	bagian := []string{"A", "B", "C"}
	gabung := strings.Join(bagian, " - ")
	fmt.Printf("Join dengan \" - \" = %q\n", gabung)
}

// -----------------------------------------------------------------------------
// 5. len(s) = jumlah BYTE; untuk karakter Unicode pakai utf8.RuneCountInString
// -----------------------------------------------------------------------------
func panjangString() {
	s := "Hello"
	fmt.Printf("%q → len() = %d (byte)\n", s, len(s))

	// Karakter non-ASCII bisa lebih dari 1 byte
	nama := "Budi"
	fmt.Printf("%q → len() = %d byte, rune count = %d\n", nama, len(nama), utf8.RuneCountInString(nama))

	emoji := "Go🚀"
	fmt.Printf("%q → len() = %d byte, rune count = %d\n", emoji, len(emoji), utf8.RuneCountInString(emoji))
}

// -----------------------------------------------------------------------------
// 6. s[i] = byte ke-i (bukan rune). Loop dengan range = per rune (karakter)
// -----------------------------------------------------------------------------
func aksesKarakter() {
	s := "Go"
	fmt.Printf("s[0] = byte %d (karakter '%c')\n", s[0], s[0])
	fmt.Printf("s[1] = byte %d (karakter '%c')\n", s[1], s[1])

	// range over string: iterasi per rune (Unicode code point)
	fmt.Print("range over \"Hi 世界\": ")
	for i, r := range "Hi 世界" {
		fmt.Printf(" [%d]=%q", i, r)
	}
	fmt.Println()

	// Konversi string → []rune jika perlu akses indeks per karakter
	rs := []rune("世界")
	fmt.Printf("[]rune(\"世界\") = %v, len = %d\n", rs, len(rs))
}

// -----------------------------------------------------------------------------
// 7. Beberapa fungsi paket strings
// -----------------------------------------------------------------------------
func paketStrings() {
	s := "  Hello, World!  "

	fmt.Printf("Asli     : %q\n", s)
	fmt.Printf("TrimSpace: %q\n", strings.TrimSpace(s))
	fmt.Printf("ToUpper  : %q\n", strings.ToUpper(s))
	fmt.Printf("ToLower  : %q\n", strings.ToLower(s))
	fmt.Printf("Contains \"World\": %t\n", strings.Contains(s, "World"))
	fmt.Printf("HasPrefix \"  He\": %t\n", strings.HasPrefix(s, "  He"))
	fmt.Printf("HasSuffix \"!  \": %t\n", strings.HasSuffix(s, "!  "))
	fmt.Printf("Replace o→0: %q\n", strings.Replace(s, "o", "0", -1)) // -1 = semua
	fmt.Printf("Split by \",\": %q\n", strings.Split("a,b,c", ","))
	fmt.Printf("Index \"World\": %d\n", strings.Index(s, "World"))
	fmt.Printf("Count \"l\": %d\n", strings.Count(s, "l"))
}

// -----------------------------------------------------------------------------
// 8. String immutable: s[i] tidak bisa di-assign; buat string baru
// -----------------------------------------------------------------------------
func immutable() {
	s := "hello"
	// s[0] = 'H'  // error: cannot assign to s[0]
	fmt.Printf("Asli: %q. Untuk \"ubah\" buat string baru, misalnya:\n", s)
	ubah := "H" + s[1:]
	fmt.Printf("  \"H\" + s[1:] = %q\n", ubah)
}

// -----------------------------------------------------------------------------
// 9. fmt.Sprintf untuk membangun string dari format (seperti Printf ke string)
// -----------------------------------------------------------------------------
func formatString() {
	nama := "Budi"
	umur := 20
	msg := fmt.Sprintf("Nama: %s, Umur: %d tahun.", nama, umur)
	fmt.Println(msg)
}
