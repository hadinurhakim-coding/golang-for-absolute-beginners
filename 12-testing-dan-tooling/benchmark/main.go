package main

import "fmt"

func main() {
	fmt.Println("Sum(100) =", Sum(100))
	fmt.Println("Benchmark: go test -bench=. -benchmem .")
}
