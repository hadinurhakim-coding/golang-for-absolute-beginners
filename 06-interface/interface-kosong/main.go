package main

import "fmt"

func cetakApaPun(x any) {
	fmt.Printf("%T: %v\n", x, x)
}

func main() {
	cetakApaPun(42)
	cetakApaPun("teks")
	cetakApaPun([]int{1, 2})
}
