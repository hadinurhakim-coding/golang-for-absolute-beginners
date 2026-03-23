//go:build !cgo

package main

import "fmt"

func main() {
	fmt.Println("CGO nonaktif (default di banyak setup Windows).")
	fmt.Println("Untuk contoh C: set CGO_ENABLED=1, pasang C compiler, lalu build ulang.")
}
