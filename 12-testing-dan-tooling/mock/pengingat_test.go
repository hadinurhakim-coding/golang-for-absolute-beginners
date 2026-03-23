package main

import "testing"

func TestPengingat_Pesan(t *testing.T) {
	p := Pengingat{
		Now: func() string { return "mock-12:00" },
	}
	got := p.Pesan()
	want := "Sekarang: mock-12:00"
	if got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}
