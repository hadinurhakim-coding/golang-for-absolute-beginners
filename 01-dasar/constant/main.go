package main

import "fmt"

const nama = "Hadi Nurhakim"

const (
	umur int = 25
	kota     = "Jakarta"
)

const (
	Senin = iota
	Selasa
	Rabu
	Kamis
	Jumat
)

func main() {
	fmt.Println(nama, umur, kota)
	fmt.Println(Senin, Selasa, Rabu, Kamis, Jumat)
}
