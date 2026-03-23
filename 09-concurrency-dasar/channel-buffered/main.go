package main

import "fmt"

func main() {
	// Buffered: kirim tidak blok sampai buffer penuh
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch, <-ch)
}
