package main

import "fmt"

type NamaLengkap string

type ID int64

type Pengguna struct {
	ID   ID
	Nama NamaLengkap
}

func main() {
	var nama NamaLengkap = "Hadi Nurhakim"
	var id ID = 1
	u := Pengguna{ID: id, Nama: nama}

	fmt.Println(u.ID, u.Nama)
}
