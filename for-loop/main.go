// =============================================================================
// FOR LOOP DI GO
// =============================================================================
// Go hanya punya satu bentuk perulangan: for.
// Bentuk: for init; condition; post { }  atau  for condition { }  atau  for { }
// for range: iterasi slice, map, string, channel.
// =============================================================================

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("========== 1. for klasik: init; condition; post ==========")
	forKlasik()

	fmt.Println("\n========== 2. for sebagai while (hanya kondisi) ==========")
	forSebagaiWhile()

	fmt.Println("\n========== 3. for tanpa kondisi (infinite + break) ==========")
	forInfinite()

	fmt.Println("\n========== 4. for range: slice ==========")
	forRangeSlice()

	fmt.Println("\n========== 5. for range: map ==========")
	forRangeMap()

	fmt.Println("\n========== 6. for range: string (rune) ==========")
	forRangeString()

	fmt.Println("\n========== 7. break dan continue ==========")
	breakContinue()

	fmt.Println("\n========== 8. for bersarang (nested) ==========")
	forBersarang()

	fmt.Println("\n========== 9. Studi kasus: jumlah 1 sampai N ==========")
	kasusJumlahN()

	fmt.Println("\n========== 10. Studi kasus: faktorial ==========")
	kasusFaktorial()

	fmt.Println("\n========== 11. Studi kasus: cetak elemen slice ==========")
	kasusCetakSlice()

	fmt.Println("\n========== 12. Studi kasus: tabel perkalian ==========")
	kasusTabelPerkalian()

	fmt.Println("\n========== 13. Studi kasus: hitung huruf vokal ==========")
	kasusHitungVokal()
}

// -----------------------------------------------------------------------------
// 1. for init; condition; post { } — init sekali, condition tiap iterasi, post setelah body
// -----------------------------------------------------------------------------
func forKlasik() {
	for i := 0; i < 5; i++ {
		fmt.Printf("  i = %d\n", i)
	}
}

// -----------------------------------------------------------------------------
// 2. for condition { } — seperti while di bahasa lain
// -----------------------------------------------------------------------------
func forSebagaiWhile() {
	n := 0
	for n < 3 {
		fmt.Printf("  n = %d\n", n)
		n++
	}
}

// -----------------------------------------------------------------------------
// 3. for { } — infinite loop; keluar pakai break
// -----------------------------------------------------------------------------
func forInfinite() {
	hitung := 0
	for {
		hitung++
		if hitung > 3 {
			break
		}
		fmt.Printf("  hitung = %d\n", hitung)
	}
	fmt.Println("  Keluar setelah break.")
}

// -----------------------------------------------------------------------------
// 4. for i, v := range slice — i = indeks, v = nilai. Pakai _ untuk skip indeks/nilai
// -----------------------------------------------------------------------------
func forRangeSlice() {
	nums := []int{10, 20, 30}
	for i, v := range nums {
		fmt.Printf("  indeks %d: nilai %d\n", i, v)
	}
	fmt.Print("  Hanya nilai: ")
	for _, v := range nums {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

// -----------------------------------------------------------------------------
// 5. for k, v := range map — k = key, v = value. Urutan tidak tetap.
// -----------------------------------------------------------------------------
func forRangeMap() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Printf("  %s => %d\n", k, v)
	}
}

// -----------------------------------------------------------------------------
// 6. for i, r := range string — i = byte index, r = rune (karakter Unicode)
// -----------------------------------------------------------------------------
func forRangeString() {
	s := "Go"
	for i, r := range s {
		fmt.Printf("  indeks %d: rune %q\n", i, r)
	}
}

// -----------------------------------------------------------------------------
// 7. break = keluar loop; continue = loncat ke iterasi berikutnya
// -----------------------------------------------------------------------------
func breakContinue() {
	fmt.Print("  Loop dengan continue (cetak ganjil saja): ")
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	fmt.Print("  Break saat i=4: ")
	for i := 0; i < 6; i++ {
		if i == 4 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// -----------------------------------------------------------------------------
// 8. for di dalam for (nested) — misalnya matriks atau pola
// -----------------------------------------------------------------------------
func forBersarang() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("  (%d,%d) ", i, j)
		}
		fmt.Println()
	}
}

// -----------------------------------------------------------------------------
// 9. Studi kasus: jumlah 1 + 2 + ... + N
// -----------------------------------------------------------------------------
func kasusJumlahN() {
	N := 5
	sum := 0
	for i := 1; i <= N; i++ {
		sum += i
	}
	fmt.Printf("  1 + 2 + ... + %d = %d\n", N, sum)
}

// -----------------------------------------------------------------------------
// 10. Studi kasus: faktorial N (N!)
// -----------------------------------------------------------------------------
func kasusFaktorial() {
	N := 5
	fak := 1
	for i := 1; i <= N; i++ {
		fak *= i
	}
	fmt.Printf("  %d! = %d\n", N, fak)
}

// -----------------------------------------------------------------------------
// 11. Studi kasus: cetak setiap elemen slice dengan nomor urut
// -----------------------------------------------------------------------------
func kasusCetakSlice() {
	buah := []string{"Apel", "Pisang", "Jeruk"}
	fmt.Println("  Daftar buah:")
	for i, b := range buah {
		fmt.Printf("    %d. %s\n", i+1, b)
	}
}

// -----------------------------------------------------------------------------
// 12. Studi kasus: tabel perkalian (misalnya 5x5)
// -----------------------------------------------------------------------------
func kasusTabelPerkalian() {
	n := 5
	fmt.Printf("  Tabel perkalian %d x %d:\n", n, n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("  %2d ", i*j)
		}
		fmt.Println()
	}
}

// -----------------------------------------------------------------------------
// 13. Studi kasus: hitung jumlah huruf vokal dalam string
// -----------------------------------------------------------------------------
func kasusHitungVokal() {
	teks := "Belajar Go"
	vokal := "aeiouAEIOU"
	count := 0
	for _, r := range teks {
		if strings.ContainsRune(vokal, r) {
			count++
		}
	}
	fmt.Printf("  Teks: %q → jumlah vokal: %d\n", teks, count)
}
