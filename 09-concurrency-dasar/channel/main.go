package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "dari goroutine"
	}()

	msg := <-ch
	fmt.Println(msg)
}
