package main

import (
	"fmt"
	"time"
)

func cetak(nama string) {
	for i := 0; i < 3; i++ {
		fmt.Println(nama, i)
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	go cetak("A")
	go cetak("B")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("selesai")
}
