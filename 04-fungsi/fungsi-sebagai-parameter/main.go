package main

import "fmt"

func terapkan(n int, fn func(int) int) int {
	return fn(n)
}

func main() {
	kuadrat := func(x int) int { return x * x }
	fmt.Println(terapkan(4, kuadrat))

	fmt.Println(terapkan(10, func(x int) int { return x + 1 }))
}
