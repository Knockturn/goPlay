package main

// Course part start: https://www.youtube.com/watch?v=YS4e4q9oBaU&t=8240s
/*
AGENDA:
	Maps
		What are they?
		Creating
		Manipulation
	Structs
		What are they?
		Creating
		Naming Convention
		Embedding
		Tags
*/

import (
	// "C:/Repos/goPlay/MapsAndStructs/Structs"
	"fmt"
)

func main() {
	fmt.Print("Maps & Structs\n")
	// Maps = Key Value pair.
	statePopulations := map[string]int{
		"California": 39250017,
		"Texas" : 27862596,
		"Florida": 20612439,
	}
	fmt.Println(statePopulations)
	/*
	NOTE: A slice cannot be a key to a map. It has to be an array
	*/
	a := map[[3]int]string{} // Note the 3 to indicate capacity making this an array not a slice which would be: []int
	fmt.Println(a)

	// Another way to initiate a map (with make)
	statePopulations2 := make(map[string]int) // Note: Will take a 3rd argument - but unsure of what it does - where in the array/slice it would be capacity.
	statePopulations2 = map[string]int{
		"California": 39250017,
		"Texas" : 27862596,
		"Florida": 20612439,
	}
	fmt.Println(statePopulations2)

	//Extracting single value
	fmt.Println(statePopulations2["Texas"])

	//Add value to map
	statePopulations2["Georgia"] = 10310371 // The return order of the data in the map is not guarenteed
	fmt.Println(statePopulations2)

	//Delete from map
	delete(statePopulations2, "Georgia")
	fmt.Println(statePopulations2)

	// Calling an element not present will return value 0 (meaning false)
	fmt.Println(statePopulations2["Georgia"])
	pop, ok := statePopulations2["Georgia"] // , ok (comman syntax for this operation but not required to be called ok). Will return true or false if the value is there or not.
	fmt.Println(pop, ok)
	_, okay := statePopulations2["Georgia"] // _ if you want to throw away the value -> note okay here because I cannot have 2 ok
	fmt.Println(okay)

	// maps share data like slices, so manipulating one will affect the original because it is the same data
	sp := statePopulations2
	delete(sp, "Texas")
	fmt.Println(sp)
	fmt.Println(statePopulations2)

	structsExamples()
}

type Doctor struct { // Capital to export outside this package (currently it would not see any fieldNames though - then you would have to capitalze those)
	number int
	actorName string
	companions []string
	episodes []string // Added to show that you dont need to add all the data. However if you use the outcommented method (Positional Syntax) it WILL break as it cannot map the data.
}

func structsExamples(){
	// Last collection type: STRUCT (Might not seem like one to begin with)
	/*
	NOTES:
	Strengh of a struct is that you have NO constrainst on the type of data within one.
	Structs can even contain other structs
	*/
	fmt.Println("\nSTRUCT examples")

	aDoctor := Doctor{
		number: 3,
		actorName: "John Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])

	// You do not have to use the name -> Positional Syntax. This is not best practice as it can become a maintenance problem. (adding extra field to the struct will then make it difficult for you to properly map the data)
	// bDoctor := Doctor{
	// 	3,
	// 	"John Pertwee",
	// 	[]string{
	// 		"Liz Shaw",
	// 		"Jo Grant",
	// 	},
	// }
	// fmt.Println(bDoctor)
	// fmt.Println(bDoctor.actorName)
	// fmt.Println(bDoctor.companions[1])

	// Another way to initialzie a struct (Following should only be lived if Data is short-lived in your app)
	cDoctor := struct{name string}{name: "John Pertwee"} // First brackets initialize the struct and its fields. 2nd brackets provide the data
	fmt.Println(cDoctor)
	anotherDoctor := cDoctor
	anotherDoctor.name = "Tom Baker"
	fmt.Println(cDoctor)
	fmt.Println(anotherDoctor)

	//Embedding
	println("\nEmbedding:")
	}

	// PROGRESS: https://youtu.be/YS4e4q9oBaU?t=9465