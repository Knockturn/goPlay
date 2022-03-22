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
			for i := range ch { // Alternative to slices. Looping over a channel range give you the first value to begin with. (in a slice you would get the index first then second time the value.)
				fmt.Println(i)
			}
			wg.Done()
		}(ch)

		go func(ch chan<- int) { // Sender/Writer only
			ch <- 42
			ch <- 27
			close(ch) // Above loop will fail unless you close the channel. Closing it will let the reciever know that no more data is coming and will prevent a deadlock in the for loop.
			// Be carefull with closing channels. You really need to mean it.
			//ch <- 16 // This will panic the app. You cannot re-open the channel. So be sure when you want to close the channel. The recover function go get around it here in the sender/writer
			wg.Done()
		}(ch)
	wg.Wait()
}

/*
	This behavior is specific to channels. It takes the bidirectional func and makes it one 
	direction by "casting" it to that behavior implicitly. This is not normal in go but special to channels.

	Buffered goroutine channels are designed to work with bursts of data. Ex:
	If you measure earthquakes. You gather 1 hours worth of data in the channel and then send it (after that 1 hour).
	Then the reader will accept 1 hours worth of data in the channel and read/take out that data - and so forth.
*/