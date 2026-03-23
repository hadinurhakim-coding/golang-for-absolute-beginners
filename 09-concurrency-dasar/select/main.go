package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "satu"
	}()
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch2 <- "dua"
	}()

	for i := 0; i < 2; i++ {
		select {
		case m := <-ch1:
			fmt.Println("dari ch1:", m)
		case m := <-ch2:
			fmt.Println("dari ch2:", m)
		case <-time.After(300 * time.Millisecond):
			fmt.Println("timeout")
		}
	}
}
