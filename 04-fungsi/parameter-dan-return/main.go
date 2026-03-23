package main

import "fmt"

func jumlah(a, b int) int {
	return a + b
}

func perkenalan(nama string, umur int) string {
	return fmt.Sprintf("%s berumur %d tahun", nama, umur)
}

func main() {
	fmt.Println(jumlah(3, 4))
	fmt.Println(perkenalan("Citra", 22))
}
