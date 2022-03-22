package main

import (
	"fmt"
	"sync"
)

// YouTube part start: https://www.youtube.com/watch?v=YS4e4q9oBaU&t=21910s

/*
Channels can be used to pass data between different goroutines
in a safe way to avoid: race conditions, memory sharing problems.
AGENDA:
	- Channel basics
	- Restricing data flow
	- Buffered channels
	- Closing channels
	- For...range loops with channels
	- Select statements (like switch statement but built for channels)
*/

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("Channels:")
	ch := make(chan int) // Strongly typed channel
	wg.Add(2)

	go func() { // Recieving goroutine
		i := <-ch // Pointing away from the channel because it pulls data FROM the channel
		fmt.Println(i)
		wg.Done()
	}()

	go func() { // Sending data to channel
		nr := 42
		ch <- nr // Arrow pointing to the way the data should flow (into the channel)
		nr = 16  // Not a pointer so it will not affect that we change it now. The channel still recieves the copy of the data from when we set it (42)
		wg.Done()
	}()

	wg.Wait()
}

// NOTE: Always have a way to close your goroutine properly