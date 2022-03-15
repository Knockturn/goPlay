package main

import (
	"fmt"
)

// Start: https://youtu.be/YS4e4q9oBaU?t=18373

func main() {
	fmt.Println("Interfaces2:")
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}
