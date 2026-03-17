package main

import "fmt"

func main() {
	umur := 25
	sudahLogin := true
	roleAdmin := false

	bolehAkses := sudahLogin && umur >= 18
	bolehAdmin := sudahLogin && roleAdmin
	bolehLewat := roleAdmin || umur >= 30
	bukanLogin := !sudahLogin

	fmt.Println(bolehAkses)
	fmt.Println(bolehAdmin)
	fmt.Println(bolehLewat)
	fmt.Println(bukanLogin)

	nama := "Hadi Nurhakim"
	fmt.Println(sudahLogin && nama == "Hadi Nurhakim")
}
