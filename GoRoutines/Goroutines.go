package main

import (
	"fmt"
	"sync"
	"time"
)

// YouTube Start: https://www.youtube.com/watch?v=YS4e4q9oBaU&t=20037s

/*
AGENDA (concurrent and parallel programming):
	Creating goroutines
	Synchronization (multiple goroutines together)
		WaitGroups
		Mutexes
	Parallelism
	Best practices
*/

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("Goroutines:")
	// Make it a goRoutine by adding go
	go sayHello()                      // Spinoff a green thread and run the function in the thread
	time.Sleep(100 * time.Millisecond) // Horrible practice but will give time for the sayHello to execute
	/*
		Traditionally a thread is large and can be expensive. Usually want to be avoided in C#
		In Go: abstraction of a thread (light). Advantage: Can start with small stack spaces. Cheap to create and destroy
		In say C# you use OS (operating system) threads and you cannot use a lot of those. But with go you could have 10s of thousands
		of goRoutines on these green threads.
	*/

	// Below is an example of why you should be careful to have reliance on vars in the callstack
	// This is called a: race condition
	var msg = "Hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye" // Because it executes after it will have changed the value to this. So be careful.
	time.Sleep(100 * time.Millisecond)
	// A way around below (pass in the arguments - this will decouple the var):
	var msg2 = "Hello2"
	go func(msg2 string) {
		fmt.Println(msg2)
	}(msg2)
	msg2 = "Goodbye2"
	time.Sleep(100 * time.Millisecond)

	// Below is how to proper wait for a goroutine (using waitgroups)
	var msg3 = "Hello3"
	wg.Add(1)
	go func(msg3 string) {
		fmt.Println(msg3)
		wg.Done()
	}(msg3)
	msg3 = "Goodbye3"
	wg.Wait()
}

func sayHello() {
	fmt.Println("Hello")
}
