package main

import "fmt"

type PersegiPanjang struct {
	Lebar, Tinggi float64
}

// Luas memakai value receiver: tidak mengubah struct; ukuran kecil.
func (p PersegiPanjang) Luas() float64 {
	return p.Lebar * p.Tinggi
}

func main() {
	pp := PersegiPanjang{Lebar: 3, Tinggi: 4}
	fmt.Println("luas:", pp.Luas())
}
