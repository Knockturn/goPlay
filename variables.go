package main

import(
	"fmt"
)

var h float32 = 43 // Declaring on package level requires this full setup

var( // declaring multiple vars can be done like this. You could group them in multiple var blocks if you want.
	actorName string = "Elisabeth Sladen"
	compantion string = "Sarah Jane Smith"
	doctorNumber int = 3
	season int = 11
)

func main(){
	fmt.println(h)
	var h float32 = 44 // You can declare a variable once pr. scope. Aka - shadowing
	fmt.println(h)

	var i int
	i = 42
	var j int = 27
	k := 99

	fmt.Printf("\n%v, %T", h, h)
	fmt.Printf("\n%v, %T", i, i)
	fmt.Printf("\n%v, %T", j, j)
	fmt.Printf("\n%v, %T", k, k)
}


// PROGRESS: https://youtu.be/YS4e4q9oBaU?t=2702