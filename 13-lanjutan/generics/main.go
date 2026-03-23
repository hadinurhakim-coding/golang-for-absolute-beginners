package main

import "fmt"

func MapSlice[T any, R any](in []T, fn func(T) R) []R {
	out := make([]R, len(in))
	for i, v := range in {
		out[i] = fn(v)
	}
	return out
}

func main() {
	angka := []int{1, 2, 3}
	kaliDua := MapSlice(angka, func(n int) int { return n * 2 })
	fmt.Println(kaliDua)
}
