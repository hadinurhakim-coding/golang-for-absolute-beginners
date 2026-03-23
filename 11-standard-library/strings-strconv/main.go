package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := strings.TrimPrefix("user:andi", "user:")
	fmt.Println("trim prefix:", s)

	b, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println("bool:", b)
	}
	f, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println("float:", f)
}
