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
	for i := 0; i < 5; i = i + 2 {
		fmt.Println(i)
	}
	for i := 0; i < 5; {
		fmt.Println(i)
		i++ // incrementing inside loop
	}

	// Same as "while" loop below
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++ // incrementing inside loop
	}

	// When to go until a special logic has been achieved inside loop (using break)
	j := 0
	for {
		fmt.Println(j)
		j++

		if j == 5 {
			break
		}
	}

	// Initializing 2 vars in a for loop
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	/*
	NOTES:
	Best practice is to not mess with the counter inside the loop
	*/
}


// PROGRESS: https://youtu.be/YS4e4q9oBaU?t=12685