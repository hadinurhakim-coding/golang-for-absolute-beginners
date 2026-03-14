// =============================================================================
// ARRAY DAN SLICE DI GO
// =============================================================================
// ARRAY: ukuran tetap [n]T. Nilai (copy saat assign/pass). Ukuran bagian dari tipe.
// SLICE: referensi ke segmen array. []T. Dinamis (append), len & cap.
// Praktik: slice dipakai hampir selalu; array untuk fixed size (buffer, matriks).
// =============================================================================

package main

import "fmt"

func main() {
	fmt.Println("========== 1. ARRAY: dasar & literal ==========")
	arrayDasar()

	fmt.Println("\n========== 2. ARRAY: ukuran tetap, zero value ==========")
	arrayFixedSize()

	fmt.Println("\n========== 3. SLICE: dasar, make(), len, cap ==========")
	sliceDasar()

	fmt.Println("\n========== 4. SLICE: append & potongan (slicing) ==========")
	sliceAppendSlicing()

	fmt.Println("\n========== 5. Perbedaan array vs slice ==========")
	perbedaanArraySlice()

	fmt.Println("\n========== 6. Studi kasus: daftar nilai siswa ==========")
	kasusNilaiSiswa()

	fmt.Println("\n========== 7. Studi kasus: riwayat suhu (slice dinamis) ==========")
	kasusRiwayatSuhu()

	fmt.Println("\n========== 8. Studi kasus: filter & agregasi ==========")
	kasusFilterAgregasi()

	fmt.Println("\n========== 9. Studi kasus: slice sebagai view array ==========")
	kasusSliceView()

	fmt.Println("\n========== 10. Studi kasus: daftar belanja (slice dinamis) ==========")
	kasusDaftarBelanja()
}

// -----------------------------------------------------------------------------
// 1. Array: deklarasi [n]T, literal [...]T atau [n]T{...}
// -----------------------------------------------------------------------------
func arrayDasar() {
	var arr [3]int           // [0 0 0], ukuran 3
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	fmt.Println("arr:", arr)

	literal := [4]int{1, 2, 3, 4}
	fmt.Println("literal [4]int:", literal)

	// [...] = compiler hitung ukuran dari jumlah elemen
	auto := [...]int{10, 20, 30}
	fmt.Println("auto [...]int:", auto, "len =", len(auto))
}

// -----------------------------------------------------------------------------
// 2. Ukuran array bagian dari tipe; [3]int dan [4]int beda tipe. Zero value = nol.
// -----------------------------------------------------------------------------
func arrayFixedSize() {
	var a [3]int
	fmt.Println("Zero value [3]int:", a)

	// a = [4]int{}  // Error: cannot use [4]int as [3]int
	b := [3]int{1, 2, 3}
	fmt.Println("b:", b, "len:", len(b))
}

// -----------------------------------------------------------------------------
// 3. Slice: []T. make([]T, len, cap), len(), cap(). Indeks 0..len-1.
// -----------------------------------------------------------------------------
func sliceDasar() {
	s := make([]int, 3, 5) // len=3, cap=5 (3 elemen, ruang 5)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	fmt.Printf("slice: %v, len=%d, cap=%d\n", s, len(s), cap(s))

	// Literal slice (tanpa ukuran di [] = slice)
	lit := []int{10, 20, 30}
	fmt.Println("literal slice:", lit)
}

// -----------------------------------------------------------------------------
// 4. append(slice, elemen...) menambah ke ujung; bisa reallocate. Slicing: s[low:high].
// -----------------------------------------------------------------------------
func sliceAppendSlicing() {
	s := []int{1, 2, 3}
	s = append(s, 4)
	fmt.Println("setelah append(s, 4):", s)

	// Potongan: s[low:high] → indeks low sampai high-1. s[1:3] = elemen indeks 1 dan 2
	potong := s[1:3]
	fmt.Println("s[1:3] =", potong)

	// append bisa banyak sekaligus
	s = append(s, 5, 6)
	fmt.Println("append(s, 5, 6):", s)
}

// -----------------------------------------------------------------------------
// 5. Array: copy by value. Slice: referensi ke backing array (share storage).
// -----------------------------------------------------------------------------
func perbedaanArraySlice() {
	arr := [2]int{1, 2}
	arr2 := arr
	arr2[0] = 99
	fmt.Println("Array: arr =", arr, "arr2 =", arr2, "(arr tidak berubah)")

	sl := []int{1, 2}
	sl2 := sl
	sl2[0] = 99
	fmt.Println("Slice: sl =", sl, "sl2 =", sl2, "(sl ikut berubah, share backing array)")
}

// -----------------------------------------------------------------------------
// 6. Studi kasus: daftar nilai — array fixed (nilai ujian N soal)
// -----------------------------------------------------------------------------
func kasusNilaiSiswa() {
	// Array: nilai per soal (misalnya 10 soal)
	nilaiPerSoal := [10]int{8, 7, 9, 6, 10, 8, 7, 9, 8, 7}
	total := 0
	for i := 0; i < len(nilaiPerSoal); i++ {
		total += nilaiPerSoal[i]
	}
	rata := float64(total) / float64(len(nilaiPerSoal))
	fmt.Printf("Nilai per soal: %v\n", nilaiPerSoal)
	fmt.Printf("Total: %d, Rata-rata: %.2f\n", total, rata)
}

// -----------------------------------------------------------------------------
// 7. Studi kasus: riwayat suhu — slice dinamis (append tiap pengukuran)
// -----------------------------------------------------------------------------
func kasusRiwayatSuhu() {
	var suhu []float64
	suhu = append(suhu, 36.5)
	suhu = append(suhu, 36.8, 37.0, 36.2)
	fmt.Println("Riwayat suhu:", suhu)

	// Rata-rata
	jum := 0.0
	for _, v := range suhu {
		jum += v
	}
	fmt.Printf("Rata-rata suhu: %.2f\n", jum/float64(len(suhu)))
}

// -----------------------------------------------------------------------------
// 8. Studi kasus: filter (nilai >= 60) dan agregasi (jumlah, banyak)
// -----------------------------------------------------------------------------
func kasusFilterAgregasi() {
	nilai := []int{55, 70, 45, 80, 65, 90}
	var lulus []int
	for _, n := range nilai {
		if n >= 60 {
			lulus = append(lulus, n)
		}
	}
	fmt.Println("Semua nilai:", nilai)
	fmt.Println("Yang lulus (>= 60):", lulus)
	fmt.Printf("Jumlah lulus: %d dari %d\n", len(lulus), len(nilai))
}

// -----------------------------------------------------------------------------
// 9. Slice tidak "memiliki" data; slice menunjuk ke bagian array. Ubah slice = ubah array.
// -----------------------------------------------------------------------------
func kasusSliceView() {
	arr := [5]int{10, 20, 30, 40, 50}
	bagian := arr[1:4] // indeks 1, 2, 3 → [20 30 40]
	fmt.Println("Array asli:", arr)
	fmt.Println("Slice arr[1:4]:", bagian)
	bagian[0] = 99 // ubah melalui slice
	fmt.Println("Setelah bagian[0]=99, array asli:", arr, "(elemen array ikut berubah)")
}

// -----------------------------------------------------------------------------
// 10. Daftar belanja: slice string yang terus di-append; mirip list dinamis.
// -----------------------------------------------------------------------------
func kasusDaftarBelanja() {
	daftar := []string{"Susu", "Roti"}
	daftar = append(daftar, "Telur")
	daftar = append(daftar, "Minyak", "Gula")
	fmt.Println("Daftar belanja:", daftar)
	fmt.Println("Jumlah item:", len(daftar))
	for i, item := range daftar {
		fmt.Printf("  %d. %s\n", i+1, item)
	}
}
