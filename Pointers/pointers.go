package main

// YouTube part start: https://www.youtube.com/watch?v=YS4e4q9oBaU&t=14637s

/*
AGENDA:
	Creating pointers
		Dereferencing pointers (using pointers to get to some underlying data)
		The 'new' function
		Working with nil
		Types with internal pointers
*/

import (
	"fmt"
)

func main() {
	fmt.Println("POINTERS:")
	a := 42
	b := a // A copy of a at point of creation. NOT pointing to the same data (like slices)
	fmt.Println(a, b)
}

// progress: https://youtu.be/YS4e4q9oBaU?t=14671
