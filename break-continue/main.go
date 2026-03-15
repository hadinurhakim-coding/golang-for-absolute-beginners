// =============================================================================
// BREAK DAN CONTINUE DI GO
// =============================================================================
// break   = keluar dari loop (for) atau switch; eksekusi lanjut setelah blok.
// continue = lewati sisa body loop, lanjut ke iterasi berikutnya.
// Label: break/continue bisa pakai label untuk keluar/lanjut ke loop terluar.
// =============================================================================

package main

import "fmt"

func main() {
	fmt.Println("========== 1. break: keluar dari for ==========")
	breakDasar()

	fmt.Println("\n========== 2. continue: loncat ke iterasi berikutnya ==========")
	continueDasar()

	fmt.Println("\n========== 3. break di switch ==========")
	breakDiSwitch()

	fmt.Println("\n========== 4. break di loop bersarang (hanya keluar loop dalam) ==========")
	breakNested()

	fmt.Println("\n========== 5. Label: break ke loop terluar ==========")
	breakLabel()

	fmt.Println("\n========== 6. Studi kasus: cari nilai lalu berhenti ==========")
	kasusCariNilai()

	fmt.Println("\n========== 7. Studi kasus: cetak hanya bilangan ganjil ==========")
	kasusCetakGanjil()

	fmt.Println("\n========== 8. Studi kasus: cari di matriks, keluar saat ketemu ==========")
	kasusCariDiMatriks()

	fmt.Println("\n========== 9. Studi kasus: skip nilai invalid (continue) ==========")
	kasusSkipInvalid()

	fmt.Println("\n========== 10. Studi kasus: loop sampai input valid ==========")
	kasusInputValid()
}

// -----------------------------------------------------------------------------
// 1. break menghentikan loop; eksekusi lanjut setelah for
// Output:
//   i = 0
//   i = 1
//   i = 2
//   i = 3
//   break saat i = 4
//   Selesai setelah loop.
// -----------------------------------------------------------------------------
func breakDasar() {
	for i := 0; i < 10; i++ {
		if i == 4 {
			fmt.Println("  break saat i = 4")
			break
		}
		fmt.Printf("  i = %d\n", i)
	}
	fmt.Println("  Selesai setelah loop.")
}

// -----------------------------------------------------------------------------
// 2. continue: tidak jalankan sisa body, lanjut ke iterasi berikutnya
// Output:
//   Cetak hanya 1,3,5,7,9: 1 3 5 7 9
// -----------------------------------------------------------------------------
func continueDasar() {
	fmt.Print("  Cetak hanya 1,3,5,7,9: ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// -----------------------------------------------------------------------------
// 3. Di switch, break hanya keluar dari switch (bukan dari for jika ada).
//    Di Go case tidak fall-through, jadi break di switch jarang dipakai.
// Output:
//   default i=0
//   case 1 (setelah ini keluar switch, for lanjut)
//   default i=2
// -----------------------------------------------------------------------------
func breakDiSwitch() {
	for i := 0; i < 3; i++ {
		switch i {
		case 1:
			fmt.Println("  case 1 (setelah ini keluar switch, for lanjut)")
		default:
			fmt.Printf("  default i=%d\n", i)
		}
	}
}

// -----------------------------------------------------------------------------
// 4. break di loop dalam hanya keluar dari loop dalam
// Output:
//   (0,0)   (0,1)   break inner saat (i,j)=(0,2)
//
//   (1,0)   (1,1)   break inner saat (i,j)=(1,2)
// -----------------------------------------------------------------------------
func breakNested() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 {
				fmt.Printf("  break inner saat (i,j)=(%d,%d)\n", i, j)
				break
			}
			fmt.Printf("  (%d,%d) ", i, j)
		}
		fmt.Println()
	}
}

// -----------------------------------------------------------------------------
// 5. Label memungkinkan break keluar ke loop terluar
// Output:
//   (0,0)   (0,1)   (0,2)
//
//   (1,0)   break Outer saat (1,1)
//   Selesai (keluar dari kedua loop).
// -----------------------------------------------------------------------------
func breakLabel() {
Outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("  break Outer saat (1,1)")
				break Outer
			}
			fmt.Printf("  (%d,%d) ", i, j)
		}
		fmt.Println()
	}
	fmt.Println("  Selesai (keluar dari kedua loop).")
}

// -----------------------------------------------------------------------------
// 6. Studi kasus: cari angka dalam slice, berhenti saat ketemu
// Output:
//   Nilai 30 ketemu di indeks 2
// -----------------------------------------------------------------------------
func kasusCariNilai() {
	nums := []int{10, 20, 30, 40, 50}
	cari := 30
	ketemu := false
	for i, n := range nums {
		if n == cari {
			fmt.Printf("  Nilai %d ketemu di indeks %d\n", cari, i)
			ketemu = true
			break
		}
	}
	if !ketemu {
		fmt.Printf("  Nilai %d tidak ditemukan.\n", cari)
	}
}

// -----------------------------------------------------------------------------
// 7. Studi kasus: cetak hanya bilangan ganjil (continue untuk genap)
// Output:
//   Bilangan ganjil 1-10: 1 3 5 7 9
// -----------------------------------------------------------------------------
func kasusCetakGanjil() {
	fmt.Print("  Bilangan ganjil 1-10: ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

// -----------------------------------------------------------------------------
// 8. Studi kasus: cari nilai di matriks, break saat ketemu
// Output:
//   Nilai 5 ketemu di (1,1)
// -----------------------------------------------------------------------------
func kasusCariDiMatriks() {
	matriks := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	cari := 5
	ketemu := false
	for i, baris := range matriks {
		for j, val := range baris {
			if val == cari {
				fmt.Printf("  Nilai %d ketemu di (%d,%d)\n", cari, i, j)
				ketemu = true
				break
			}
		}
		if ketemu {
			break
		}
	}
}

// -----------------------------------------------------------------------------
// 9. Studi kasus: proses slice, skip nilai yang invalid (misalnya <= 0)
// Output:
//   Hanya nilai positif: 3 4 5 6
// -----------------------------------------------------------------------------
func kasusSkipInvalid() {
	nums := []int{3, -1, 4, 0, 5, -2, 6}
	fmt.Print("  Hanya nilai positif: ")
	for _, n := range nums {
		if n <= 0 {
			continue
		}
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

// -----------------------------------------------------------------------------
// 10. Studi kasus: simulasi loop sampai dapat input valid (break saat valid)
// Output:
//   Input 0 invalid, coba lagi.
//   Input -1 invalid, coba lagi.
//   Input valid: 3. Keluar loop.
// -----------------------------------------------------------------------------
func kasusInputValid() {
	// Simulasi: anggap "valid" jika nilai dalam 1-5
	nilaiCoba := []int{0, -1, 3} // coba 0, -1, lalu 3 (valid)
	for _, v := range nilaiCoba {
		if v < 1 || v > 5 {
			fmt.Printf("  Input %d invalid, coba lagi.\n", v)
			continue
		}
		fmt.Printf("  Input valid: %d. Keluar loop.\n", v)
		break
	}
}
