package main

import "fmt"

func main() {
	// Slice: referensi ke segmen array di balik layar; ukuran dinamis.
	buah := []string{"apel", "jeruk"}
	fmt.Println("literal:", buah, "len:", len(buah), "cap:", cap(buah))

	// make([]T, len, cap) — kapasitas awal mengurangi alokasi berulang saat append.
	angka := make([]int, 0, 4)
	angka = append(angka, 1, 2)
	angka = append(angka, 3)
	fmt.Println("append:", angka, "cap:", cap(angka))

	// append bisa menambah slice lain dengan ...
	lain := []int{4, 5}
	angka = append(angka, lain...)
	fmt.Println("append slice:", angka)

	// copy(dst, src) — menyalin elemen hingga minimum len(dst), len(src).
	dst := make([]int, 2)
	n := copy(dst, angka)
	fmt.Println("copy ke dst:", dst, "elemen tersalin:", n)

	// Irisan: s[low:high] (high eksklusif); berbagi backing dengan sumber.
	s := []int{10, 20, 30, 40, 50}
	tengah := s[1:4]
	fmt.Println("tengah:", tengah)
	tengah[0] = 200 // mengubah elemen yang sama dengan s[1]
	fmt.Println("s setelah ubah tengah[0]:", s)

	// Dari array (materi array): arr[:] menghasilkan slice ke seluruh array.
	arr := [3]int{7, 8, 9}
	v := arr[:]
	fmt.Println("slice penuh dari array:", v)
}
