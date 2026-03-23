package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int, 5)
	var wg sync.WaitGroup

	worker := func(id int) {
		defer wg.Done()
		for j := range jobs {
			fmt.Printf("worker %d proses job %d\n", id, j)
		}
	}

	const nWorker = 3
	for w := 0; w < nWorker; w++ {
		wg.Add(1)
		go worker(w)
	}

	for i := 1; i <= 7; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}
