package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	var wg sync.WaitGroup

	inisialisasi := func() {
		fmt.Println("init sekali saja")
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(inisialisasi)
		}()
	}
	wg.Wait()
}
