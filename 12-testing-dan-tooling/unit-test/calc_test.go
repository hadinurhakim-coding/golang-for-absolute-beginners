package main

import "testing"

func TestTambah(t *testing.T) {
	got := Tambah(2, 3)
	if got != 5 {
		t.Fatalf("Tambah(2,3) = %d, mau 5", got)
	}
}
