package main

import (
	"fmt"
	"strconv"
)

// Course part start: https://www.youtube.com/watch?v=YS4e4q9oBaU&t=2148s

/*
NOTES:
Declaring variable names - use prescriptibe(verbose) names the longer the lifespan of the var is. In this file because they are used immidiatly single letters is acceptable.
Acronyms should be ALL UPPER CASE: myURL (not myUrl), theHTTP (not theHttp), for readability. See example below:
Naming conventions: Pascal or camelCase (although acronym(HTTP, URL) are capitalized)
*/
var theURL string = "https://google.com"
var G string = "public" // Upper case in first letter of variable name declares it to be available globally - while lower case (like below) declares it for the local package only.
var h float32 = 43 // Declaring on package level requires this full setup

var( // declaring multiple vars can be done like this. You could group them in multiple var blocks if you want.
	actorName string = "Elisabeth Sladen"
	compantion string = "Sarah Jane Smith"
	doctorNumber int = 3
	season int = 11
)

func main(){
	fmt.Println(theURL)	
	fmt.Println(G)

	fmt.Println(h)
	var h float32 = 44 // You can declare a variable once pr. scope. Aka - shadowing
	fmt.Println(h)

	var i int
	i = 42
	var j int = 27 // This type of declaration should only be used if compiler would guesse wrong (EX: A nr that could be a float but often an integer)
	k := 99

	fmt.Printf("%v, %T\n", h, h)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)

	// Converting from 1 var type 2 another:
	fmt.Println("Converting 1 type 2 another")
	h = float32(i)
	fmt.Printf("%v, %T\n", h, h)
	// Careful: If converting a float 23.5 to int you will lose information - the .5
	
	// from int to string
	fmt.Println("From int to string:")
	G = string(i)
	fmt.Printf("%v, %T\n", G, G)
	G = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", G, G)

}