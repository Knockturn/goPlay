package main

// Youtube part: https://www.youtube.com/watch?v=YS4e4q9oBaU&t=12077s
/*
AGENDA: (control flow)
	For statements
		simple loops
		exiting early
		looping through collections
*/

import (
	"fmt"
)

func main() {
	fmt.Println("LOOPING")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}