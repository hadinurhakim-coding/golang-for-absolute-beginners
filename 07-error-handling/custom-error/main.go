package main

import (
	"errors"
	"fmt"
)

var ErrStokHabis = errors.New("stok habis")

type ValidasiError struct {
	Field string
	Msg   string
}

func (e *ValidasiError) Error() string {
	return fmt.Sprintf("validasi %s: %s", e.Field, e.Msg)
}

func ambilBarang(stok int) error {
	if stok < 0 {
		return &ValidasiError{Field: "stok", Msg: "negatif"}
	}
	if stok == 0 {
		return ErrStokHabis
	}
	return nil
}

func main() {
	err := ambilBarang(0)
	if errors.Is(err, ErrStokHabis) {
		fmt.Println("sentinel:", err)
	}

	err2 := ambilBarang(-1)
	var ve *ValidasiError
	if errors.As(err2, &ve) {
		fmt.Println("custom type:", ve.Field, ve.Msg)
	}
}
