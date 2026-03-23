package main

import "fmt"

type Counter struct {
	n int
}

// Increment memakai pointer receiver agar mutasi terlihat di pemanggil.
func (c *Counter) Increment() {
	c.n++
}

func (c *Counter) Nilai() int {
	return c.n
}

func main() {
	var c Counter
	c.Increment()
	c.Increment()
	fmt.Println(c.Nilai())

	p := &Counter{}
	p.Increment()
	fmt.Println(p.Nilai())
}
