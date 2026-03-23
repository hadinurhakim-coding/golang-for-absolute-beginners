package main

import "fmt"

type Orang struct {
	Nama string
}

func (o Orang) Perkenalan() string {
	return "Saya " + o.Nama
}

type Karyawan struct {
	Orang // embedding: field tanpa nama — method & field "dipromosikan"
	Role  string
}

func main() {
	k := Karyawan{
		Orang: Orang{Nama: "Citra"},
		Role:  "Developer",
	}
	// Akses langsung ke field/method dari Orang
	fmt.Println(k.Nama, k.Role)
	fmt.Println(k.Perkenalan())
}
