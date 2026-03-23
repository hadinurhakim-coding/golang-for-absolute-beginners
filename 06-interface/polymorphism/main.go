package main

import "fmt"

type Bentuk interface {
	Luas() float64
}

type Lingkaran struct {
	JariJari float64
}

func (l Lingkaran) Luas() float64 {
	return 3.14 * l.JariJari * l.JariJari
}

type Persegi struct {
	Sisi float64
}

func (p Persegi) Luas() float64 {
	return p.Sisi * p.Sisi
}

func totalLuas(bentuk []Bentuk) float64 {
	t := 0.0
	for _, b := range bentuk {
		t += b.Luas()
	}
	return t
}

func main() {
	shapes := []Bentuk{
		Lingkaran{JariJari: 1},
		Persegi{Sisi: 2},
	}
	fmt.Println("total luas:", totalLuas(shapes))
}
