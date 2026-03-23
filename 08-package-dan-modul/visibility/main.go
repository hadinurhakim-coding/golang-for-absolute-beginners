package main

import (
	"fmt"

	"golang-for-absolute-beginners/08-package-dan-modul/visibility/akun"
)

func main() {
	p := akun.BuatProfil("citra", "rahasia123")
	fmt.Println(p.Username)
	// p.password — tidak bisa: field tidak diekspor
	fmt.Println(p.Ringkasan())
}
