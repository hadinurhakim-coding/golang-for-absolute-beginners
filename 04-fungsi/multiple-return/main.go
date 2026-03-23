package main

import (
	"errors"
	"fmt"
)

func bagi(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("pembagi nol")
	}
	return a / b, nil
}

func main() {
	hasil, err := bagi(10, 2)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("10/2 =", hasil)

	if _, err := bagi(1, 0); err != nil {
		fmt.Println("bagi gagal:", err)
	}
}
