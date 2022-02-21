package main

//Course part start: https://www.youtube.com/watch?v=YS4e4q9oBaU&t=5189s

/*
Agenda:
Naming convention
Typed Constants
Untyped Constants
Enumerated constants
Enumeration Expressions
*/

import (
	"fmt"
)

const a string = "A"

// Consts work the same most of the way as a variable (var) - however it is immutable (unchangeable)
func main() {
	// other language: const MY_CONST
	const myConst int = 42 // Remember lower vs upper case begining if you want it local or global
	fmt.Printf("%v, %T\n", myConst, myConst)
	// const myConst2 float64 = math.Sin(1.57) // Executing function at compile time is not allowed for constants!

	const a string = "B" // Shadowing also allowed (inner declaration will win)

	const b = 16 // When declaring like this it will look for usage and might convert it from int32 to int8 - just like normal variables. 
	fmt.Printf("%v, %T\n", myConst + b, myConst + b)

	// Enumerated constants
	fmt.Println("Enumerated Constants")
	const c = iota
	fmt.Printf("%v, %T\n", c, c)
	const (
		d = iota
		e = iota
		f = iota
	)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)
	// Infer pattern of assignments
	fmt.Println("Infer pattern:")
	const ( // iota is scoped to a constant block - so notice how it resets.
		g = iota
		h
		i
	)
	fmt.Printf("%v\n", g)
	fmt.Printf("%v\n", h)
	fmt.Printf("%v\n", i)
}

// PROGRESS: https://youtu.be/YS4e4q9oBaU?t=5742