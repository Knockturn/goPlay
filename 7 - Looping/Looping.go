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
	fmt.Println("\nSame as While loop:")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++ // incrementing inside loop
	}

	// When to go until a special logic has been achieved inside loop (using break)
	fmt.Println("\nBreak:")
	j := 0
	for {
		fmt.Println(j)
		j++

		if j == 5 {
			break
		}
	}

	// Alternative for continue above is continue (continue is not used alot)
	fmt.Println("\nContinue:")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// Initializing 2 vars in a for loop
	fmt.Println("\nInitializing 2 vars")
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i, j)
	}

	// To break out of an inner loop and also break the outer loop: You can use a "label"
	fmt.Println("\nInner Loop")
	Loop:
		for i := 1; i <= 3; i++ {
			for j := 1; j <= 3; j++ { // inner loop
				fmt.Println(i * j)
				if i * j >= 3 {
					break Loop
				}
			}
		}

	// looping through arbitrary(unkown) sized slice
	fmt.Println("\nSlice")
	s := []int{1, 2, 3}
	for k, v := range s { // KEY, VALUE (useful perhaps in maps.Ex K&V: Texas 105687)
		// if you only want value you can do this: "for _, v := range s" (and of course only k for key (k is just a representation - could be named anything))
		fmt.Println(k,v)
	}

	/**/

	/*
	NOTES:
	Best practice is to not mess with the counter inside the loop.
	For loops can be used for channels as well - but require special logic: check channels tutorial
	*/
}