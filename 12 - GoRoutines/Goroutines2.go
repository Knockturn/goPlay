package main

import (
	"fmt"
	"runtime"
	"sync"
)

// YouTube Part Explaining MUTEX: https://youtu.be/YS4e4q9oBaU?t=20906
var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{} // Mutex is basically a lock that the app has to honor/follow.
/*
	Read/Write Mutex:
	As many thing as wants can READ the data, but only 1 can write. And if anything is reading
	then we cannot write to it at all. The writer will wait for the readers to be done, then lock it
	and write and then unlock when done so it can be read again.

	Mutex force the application to be executed in a single thread way.
	So mutex and goroutines doesn't mix very well in this example.
	It would run better if we remove all goroutines and mutex so it runs as normal functions
*/
func main() {
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
	runtime.GOMAXPROCS(100) // Sets maximum nr of OS threads (CPUs)
	/*
		General rule of thumb:
		1 OS thread pr CPU.
		The higher you set the more is required.
		Before releasing to production test different GOMAXPROCS sizes to see which one is the fastest/best
	*/
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock() // locks FOR reading (to be able to read)
		go sayHello()
		m.Lock() // lock to be able to write
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
