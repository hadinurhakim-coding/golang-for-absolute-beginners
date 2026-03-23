package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "  Hello, Go World  "

	fmt.Println("Contains 'Go':", strings.Contains(s, "Go"))
	fmt.Println("Split:", strings.Split("a,b,c", ","))
	fmt.Println("Join:", strings.Join([]string{"x", "y", "z"}, "-"))
	fmt.Println("Replace:", strings.ReplaceAll("foo foo", "foo", "bar"))
	fmt.Println("TrimSpace:", strings.TrimSpace(s))

	// strconv: string ↔ angka
	n, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println("Atoi error:", err)
	} else {
		fmt.Println("Atoi:", n)
	}
	fmt.Println("Itoa:", strconv.Itoa(100))

	// String vs []byte: string immutable; slice byte bisa diubah (salinan).
	teks := "halo"
	b := []byte(teks)
	b[0] = 'H'
	fmt.Println("[]byte ubah salinan:", string(b), "| string asli:", teks)

	// Per karakter Unicode: range pada string mengeluarkan rune (bukan selalu 1 byte).
	emoji := "Hi 世界"
	for i, r := range emoji {
		fmt.Printf("indeks %d rune %q\n", i, r)
	}
}
