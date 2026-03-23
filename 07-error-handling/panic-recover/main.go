package main

import "fmt"

func aman() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover menangkap panic:", r)
		}
	}()
	panic("kondisi fatal contoh")
}

func main() {
	fmt.Println("sebelum aman")
	aman()
	fmt.Println("setelah aman — program lanjut")
}
