package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := range 50 {
		wg.Go(func() {
			fmt.Printf("Processing order %d\n", i)
			time.Sleep(2 * time.Second)
			fmt.Printf("Processed order %d\n", i)
		})
	}
	wg.Wait()
}
