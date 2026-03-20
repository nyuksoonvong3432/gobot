package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 50; i++ {
		fmt.Printf("Processing order %d\n", i)
		time.Sleep(2 * time.Second)
		fmt.Printf("Processed order %d\n", i)
	}
}
