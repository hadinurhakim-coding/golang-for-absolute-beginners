package main

import (
	"errors"
	"fmt"
)

func bagi(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("pembagi tidak boleh nol")
	}
	return a / b, nil
}

func bacaID(id int) error {
	if id < 0 {
		return fmt.Errorf("id tidak valid: %d", id)
	}
	return nil
}

func main() {
	if _, err := bagi(1, 0); err != nil {
		fmt.Println("bagi:", err)
	}
	if err := bacaID(-1); err != nil {
		fmt.Println("bacaID:", err)
	}
}
