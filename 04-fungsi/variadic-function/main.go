package main

import "fmt"

func total(angka ...int) int {
	t := 0
	for _, n := range angka {
		t += n
	}
	return t
}

func main() {
	fmt.Println(total(1, 2, 3))

	slice := []int{4, 5, 6}
	fmt.Println(total(slice...))
}
