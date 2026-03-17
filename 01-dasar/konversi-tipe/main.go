package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 42
	var b float64 = float64(a)
	var c int = int(b)

	var d string = "100"
	e, err := strconv.Atoi(d)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(a, b, c, d, e)
}
