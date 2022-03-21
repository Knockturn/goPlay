package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	for j := 0; j < 5; j++ { // Creates 5 sets of goroutines each set containing 2 goroutines. A total of 10 goroutines
		wg.Add(2)
		go func() { // Reciever/Reader
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()

		go func() { // Sender/Writer
			ch <- 42
			wg.Done()
		}()
	}
	wg.Wait()
}
/*
	If you move the reciever out of the loop it would fail (deadlock). The channel is unbufferen meaning only 1 set of data can be in the channel at a time. So if nothing is there to recieve it it will fail.
*/
