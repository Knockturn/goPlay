package main

import (
	"fmt"
)

// Start: https://youtu.be/YS4e4q9oBaU?t=19116
func main() {
	fmt.Println("Type Switches interface:")
	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("I don't know what i is")
	}
}
