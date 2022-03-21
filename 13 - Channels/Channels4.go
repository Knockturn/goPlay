package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

// Unidirectional Channel

func main() {
	ch := make(chan int)
		wg.Add(2)
		go func(ch <-chan int) { // Reciever/Reader only
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}(ch)

		go func(ch chan<- int) { // Sender/Writer only
			ch <- 42
			wg.Done()
		}(ch)
	wg.Wait()
}

/*
	This behavior is specific to channels. It takes the bidirectional func and makes it one 
	direction by "casting" it to that behavior implicitly. This is not normal in go but special to channels.
*/
