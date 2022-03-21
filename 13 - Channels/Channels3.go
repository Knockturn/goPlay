package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
// Bidirectional channels (sending and recieving (can get confusing - seems best to avoid))
func main() {
	ch := make(chan int)
		wg.Add(2)
		go func() { // Reciever/Reader
			i := <-ch
			fmt.Println(i)
			ch <- 27 // Sending/Writing to below
			wg.Done()
		}()

		go func() { // Sender/Writer
			ch <- 42
			fmt.Println(<-ch) // Recieving/Reading from above
			wg.Done()
		}()
	wg.Wait()
}
