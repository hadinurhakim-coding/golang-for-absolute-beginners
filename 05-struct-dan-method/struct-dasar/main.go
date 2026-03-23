package main

import "fmt"

type Pengguna struct {
	ID   int
	Nama string
}

func main() {
	var kosong Pengguna
	fmt.Println("zero value:", kosong)

	u := Pengguna{ID: 1, Nama: "Andi"}
	fmt.Println(u, u.Nama)

	v := Pengguna{}
	v.ID = 2
	v.Nama = "Budi"
	fmt.Println(v)
}
