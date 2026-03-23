package main

import "testing"

func TestPangkatDua(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want int
	}{
		{"nol", 0, 0},
		{"tiga", 3, 9},
		{"min lima", -5, 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PangkatDua(tt.in); got != tt.want {
				t.Errorf("PangkatDua(%d) = %d, mau %d", tt.in, got, tt.want)
			}
		})
	}
}
