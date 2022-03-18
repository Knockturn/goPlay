package main

import (
	"fmt"
	"time"
)

// Best practice example:
// How to check for race condition
// Execute with -race, ex: go run -race Goroutine.go

func main() {
	var msg = "Hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}
