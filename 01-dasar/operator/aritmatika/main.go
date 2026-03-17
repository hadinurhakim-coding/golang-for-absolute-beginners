package main

import "fmt"

func main() {
	x := 10
	y := 3

	fmt.Println(x+y, x-y, x*y, x/y, x%y)

	a := 10.0
	b := 4.0
	fmt.Println(a/b, a+b, a-b, a*b)

	fmt.Println(2+3*4, (2+3)*4)
}
