package main

import "fmt"

func main() {
	// Bentuk penuh: init; kondisi; post (mirip for di C)
	fmt.Println("--- for klasik ---")
	for i := 0; i < 3; i++ {
		fmt.Println("i =", i)
	}

	// Hanya kondisi = seperti while
	fmt.Println("--- for kondisi ---")
	j := 0
	for j < 3 {
		fmt.Println("j =", j)
		j++
	}

	// Tanpa kondisi = loop tak terbatas (hentikan dengan break)
	fmt.Println("--- for tanpa kondisi + break ---")
	k := 0
	for {
		if k >= 2 {
			break
		}
		fmt.Println("k =", k)
		k++
	}

	// for range pada slice (nanti dipakai lagi untuk map, string, channel)
	fmt.Println("--- for range slice ---")
	buah := []string{"apel", "jeruk", "mangga"}
	for idx, nama := range buah {
		fmt.Println(idx, nama)
	}
	// Hanya indeks
	for i := range buah {
		fmt.Println("indeks saja:", i)
	}
	// Hanya nilai (abaikan indeks dengan _)
	for _, nama := range buah {
		fmt.Println("nama:", nama)
	}
}
