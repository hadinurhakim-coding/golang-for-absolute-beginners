package main

import "testing"

func TestSum(t *testing.T) {
	if Sum(10) != 55 {
		t.Fatal("salah")
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1000)
	}
}
