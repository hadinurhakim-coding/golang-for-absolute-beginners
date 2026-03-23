package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var n int
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			n++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("n:", n)
}
