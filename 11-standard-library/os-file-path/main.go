package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.MkdirTemp("", "belajar-go-*")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(dir)

	p := filepath.Join(dir, "data.txt")
	if err := os.WriteFile(p, []byte("isi file"), 0o644); err != nil {
		fmt.Println(err)
		return
	}
	b, err := os.ReadFile(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("baca:", string(b))

	fmt.Println("Base:", filepath.Base(p))
	if home, err := os.UserHomeDir(); err == nil {
		fmt.Println("UserHomeDir:", home)
	}
}
