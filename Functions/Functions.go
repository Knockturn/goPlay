package main

// YouTube part start: https://www.youtube.com/watch?v=YS4e4q9oBaU&t=15690s

/*
AGENDA:
	Basix syntax
	Parameters
	Return values
	Anonymous functions
	Functions as types (first class citizen)
	Method (special type of function)
*/

import (
	"fmt"
)

func main() { // This is the most basic function in go. As all of Go - capital start = global/shared
	fmt.Println("Functions:")

	sayMessage("Hello Go!")
	sayGreeting("Hello", "Phi!")
}

// Parameters
func sayMessage(msg string) {
	fmt.Println(msg)
}

func sayGreeting(greeting, name string) { // Multiple parameters of same type can be done like this
	fmt.Println(greeting, name)
}

// Progress: https://youtu.be/YS4e4q9oBaU?t=16081
