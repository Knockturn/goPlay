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

	greeting := "Hello"
	name := "Stacey"
	pointingGreeting(&greeting, &name)
	println(name)
	sum("The sum is:", 1, 2, 3, 4, 5)
	s := returnSum(1, 2, 3, 4, 5)
	println("The sum is:", s)
	sp := returnSumAsPointer(2, 4, 6)
	println("The sum is:", *sp)
	println("The sum is:", returnSumNamedValue(1, 2, 3, 4, 5))

	// Multiple return func
	println("Multiple return")
	d, err := divide_multipleReturnValues(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	// Anonymous function
	func() {
		fmt.Println("Anonymous function")
	}()

	// Another example - in general dont use anonymous func unless you need to protect a variable by having it local inside another func.
	for i := 0; i < 5; i++ {
		func(i int) { // Best practice IN CASE you use async functions - as it would otherwhise give weird behaviour(i would count wrong)
			fmt.Println(i)
		}(i)
	}

	/*
		https://youtu.be/YS4e4q9oBaU?t=17199
		You could use an anomoymous func as a variable. However it looks unreadable to me - so please avoid.
	*/

	fmt.Println("\nMethods:")
	/*
		A method is just a function that executes in a known context.
		A known context = any type (in Go)
		Ex: struct, int, string etc
	*/
	g := greeter{
		greeting: "Hello",
		name:     "Phi",
	}
	g.greet()

	g.pointerGreet()
}

// Parameters
func sayMessage(msg string) {
	fmt.Println(msg)
}

func sayGreeting(greeting, name string) { // Multiple parameters of same type can be done like this
	fmt.Println(greeting, name)
	// Changing name here will only update the local version of name. Not the underlying data
}

// Passing in a pointer is the only way to work on the same underlying data and it can also be more effecient that copying like above
func pointingGreeting(greeting, name *string) {
	// Passing in a large data structure is best use of pointers since it is way more effecient.
	// A string is "small" and it's not so hard on the application to copy it every time.
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

// Variatic parameterbad
func sum(msg string, values ...int) { // Will wrap up all the paremeters passed to this method in a int slice. Needs to be last paremter though
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}

// Return
func returnSum(values ...int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

// Return local variable as pointer
func returnSumAsPointer(values ...int) *int {
	/*
		In a lot of languages this is not a valid operation as the local var will be cleared from memory after function is done running.
		So pointing to data in here would be gone - hence invalid.
		But in Go - it promotes the local var to heap memory (shared memory) once you make it a pointer (local memory: stack memory).
	*/
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

// Named return (not often used)
// Can be confusing to read and in long functions you have no idea what you are returning possibly. So be careful
func returnSumNamedValue(values ...int) (result int) {
	for _, v := range values {
		result += v
	}
	return
}

// Multiple return values
func divide_multipleReturnValues(a, b float64) (float64, error) { // You can return as many values as you want
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil // Returning nil if no error
}

// Methods: https://youtu.be/YS4e4q9oBaU?t=17256
type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() { // (g greeter) -> that part is what makes this func into a method!
	// Gets a copy of the greeter object and its assigned to the variable name g in this context
	fmt.Println(g.greeting, g.name)
	// Since it is a copy of the object sent in - you will not change the underlying data by changing in here.
	// You can however change it to a pointer reciever and it will then be the underlying data
}

func (g *greeter) pointerGreet() {
	fmt.Println(g.greeting, g.name)
}
