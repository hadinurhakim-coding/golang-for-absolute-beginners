package main

import (
	"errors"
	"fmt"
)

type detailError struct{ kode int }

func (e *detailError) Error() string {
	return fmt.Sprintf("kode %d", e.kode)
}

func langkahA() error {
	return fmt.Errorf("A gagal: %w", errors.New("koneksi timeout"))
}

func langkahB() error {
	return errors.New("B: data tidak ada")
}

func jalankan() error {
	errA := langkahA()
	errB := langkahB()
	// Gabung beberapa error (Go 1.20+)
	return errors.Join(errA, errB)
}

func main() {
	err := jalankan()
	fmt.Println("joined:", err)

	// Unwrap rantai %w + errors.Is
	dasar := errors.New("akar")
	wrapped := fmt.Errorf("lapisan: %w", dasar)
	fmt.Println("Is akar:", errors.Is(wrapped, dasar))

	inner := &detailError{kode: 404}
	w2 := fmt.Errorf("request: %w", inner)
	var d *detailError
	if errors.As(w2, &d) {
		fmt.Println("As *detailError:", d.kode)
	}
}
