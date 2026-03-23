package main

import (
	"fmt"
	"time"
)

func main() {
	sekarang := time.Now()
	fmt.Println("Now:", sekarang)

	d := 2*time.Hour + 30*time.Minute
	fmt.Println("Duration:", d)

	t, err := time.Parse(time.RFC3339, "2025-01-15T10:00:00Z")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Parsed:", t.Format("2006-01-02"))

	tick := time.NewTicker(100 * time.Millisecond)
	defer tick.Stop()
	<-tick.C
	fmt.Println("satu tick")
}
