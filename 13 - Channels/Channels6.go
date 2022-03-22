package main

import (
	"fmt"
	"sync"
)

/*
Alternative way of doing the same as in Channels5.go -> In this case channels5.go's way of doing it would make more sense. But just know you can do the bottom way also.
*/

// Start: https://youtu.be/YS4e4q9oBaU?t=23168
var wg = sync.WaitGroup{}

func main()  {
	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int)  {
		for {
			if i, ok := <- ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int)  {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}