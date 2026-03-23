package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	ch := make(chan string)
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch <- "lambat"
	}()

	select {
	case v := <-ch:
		fmt.Println(v)
	case <-ctx.Done():
		fmt.Println("batal:", ctx.Err())
	}
}
