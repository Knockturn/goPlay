package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()

		go func() {
			ch <- 42
			wg.Done()
		}()
	}
	wg.Wait()
}

// PROGRESS: https://youtu.be/YS4e4q9oBaU?t=22314
