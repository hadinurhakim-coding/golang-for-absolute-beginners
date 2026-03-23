package main

import "fmt"

func main() {
	var f func(string)
	f = func(s string) {
		fmt.Println("->", s)
	}
	f("satu")

	g := func(x int) int { return x * 2 }
	fmt.Println(g(5))
}
