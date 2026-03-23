package main

import "fmt"

// WaktuSekarang dipakai untuk injeksi waktu (mudah di-mock di tes).
type WaktuSekarang func() string

type Pengingat struct {
	Now WaktuSekarang
}

func (p Pengingat) Pesan() string {
	return "Sekarang: " + p.Now()
}

func main() {
	p := Pengingat{Now: func() string { return "asli" }}
	fmt.Println(p.Pesan())
}
