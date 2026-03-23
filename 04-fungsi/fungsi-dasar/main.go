package main

import "fmt"

// sapa mencetak sapaan (fungsi lokal, huruf kecil — tidak diekspor dari paket).
func sapa(nama string) {
	fmt.Println("Halo,", nama)
}

func main() {
	sapa("Andi")
	sapa("Budi")
}
