package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

/*
	Buffered channels
*/
// Start: https://youtu.be/YS4e4q9oBaU?t=22651

func main() {
	ch := make(chan int, 50) // Added a buffer (, 50) - this gives the channel an internal data storage of 50 ints.
		wg.Add(2)
		go func(ch <-chan int) { // Reciever/Reader only
			i := <-ch
			fmt.Println(i)

			// Below is added to read the 2nd int in the channel
			i = <-ch
			fmt.Println(i)
			wg.Done()
		}(ch)

		go func(ch chan<- int) { // Sender/Writer only
			ch <- 42
			ch <- 27
			wg.Done()
		}(ch)
	wg.Wait()
}

/*
	This behavior is specific to channels. It takes the bidirectional func and makes it one 
	direction by "casting" it to that behavior implicitly. This is not normal in go but special to channels.
*/

// Progress: https://youtu.be/YS4e4q9oBaU?t=22754