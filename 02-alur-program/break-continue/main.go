package main

import "fmt"

func main() {
	// continue: lompat ke iterasi berikutnya
	fmt.Println("--- continue (bilangan genap saja) ---")
	for i := 0; i < 6; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	// break: keluar dari loop terdekat
	fmt.Println("--- break ---")
	for i := 0; i < 10; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	// Label: break/continue ke loop luar (berguna di loop bersarang)
	fmt.Println("--- label + break luar ---")
Outer:
	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ {
			if a == 1 && b == 1 {
				break Outer
			}
			fmt.Printf("a=%d b=%d\n", a, b)
		}
	}
	fmt.Println("selesai setelah break Outer")

	fmt.Println("--- label + continue luar ---")
Loop:
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if y == 1 {
				continue Loop
			}
			fmt.Printf("x=%d y=%d\n", x, y)
		}
	}
}
