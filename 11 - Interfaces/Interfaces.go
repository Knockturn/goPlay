package main

// YouTube start part: https://www.youtube.com/watch?v=YS4e4q9oBaU&t=17879s

import (
	"fmt"
)

/*
Interfaces in Go are one of the reasons they are so scalable and maintanable.
AGENDA:
	Basics
	Composing interfaces (can make interfaces of interfaces)
	Type conversion (a bit different now with interfaces than previous tutorial)
		The empty interface
		Type switches
	Implementing with values vs. pointers
	Best practices
*/

func main() {
	fmt.Println("Interfaces")

	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}

/*
Naming Convention:
Interfaces should be named why what they do. If it is a single method one like below
then you name it what it does (write) and with 'er' at the end like an action: writer
*/
type Writer interface { // Interfaces describe behavior (not data like a struct)
	Write([]byte) (int, error) // Takes in a slice of byte and returns an int and error (if there is one)
	// above from the io package. We are recreating below for this example.
}

/*
https://youtu.be/YS4e4q9oBaU?t=18246
Interfaces can be a way to test logic without needed the specific type.
Ex:
Database and SQL package from Go is concrete methods (a struct)
Therefor for testing it without a DB you can implement an interface that replicates the signature
and the DB object fromt he sequal package will automatically implement it.
*/

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
