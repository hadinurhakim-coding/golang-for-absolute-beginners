package main

import "fmt"

func penghitung() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func main() {
	next := penghitung()
	fmt.Println(next(), next(), next())

	lain := penghitung()
	fmt.Println("mulai lagi:", lain())
}
