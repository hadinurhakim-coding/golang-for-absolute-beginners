package main

import "fmt"

type Bersuara interface {
	Suara() string
}

type Anjing struct{}

func (Anjing) Suara() string { return "guk" }

type Kucing struct{}

func (Kucing) Suara() string { return "meong" }

func dengarkan(b Bersuara) {
	fmt.Println(b.Suara())
}

func main() {
	dengarkan(Anjing{})
	dengarkan(Kucing{})
}
