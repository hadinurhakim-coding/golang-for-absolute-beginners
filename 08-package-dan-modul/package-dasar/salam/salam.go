package salam

import "fmt"

// Halo diekspor — bisa dipakai dari paket lain.
func Halo(nama string) {
	fmt.Println("Halo,", nama)
}

func kecil(nama string) {
	fmt.Println("hai", nama)
}

// PanggilKecil memakai fungsi tidak diekspor di dalam paket yang sama.
func PanggilKecil(nama string) {
	kecil(nama)
}
