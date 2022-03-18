package main

// YouTube part start: https://www.youtube.com/watch?v=YS4e4q9oBaU&t=14637s

/*
AGENDA:
	Creating pointers
		Dereferencing pointers (using pointers to get to some underlying data)
			(Complex types: ex structs, are automatically derefenced)
		The 'new' function
		Working with nil
		Types with internal pointers
*/

import (
	"fmt"
)

func main() {
	fmt.Println("POINTERS:")
	a := 42
	b := a // A copy of a at point of creation. NOT pointing to the same data (like slices)
	fmt.Println(a, b)
	a = 23
	fmt.Println(a, b)

	var c *int = &a   // This now has a pointer (&) to a (same underlying data as a now)
	fmt.Println(a, c) // Output from this ex: 0xc0000140e0 is the location in memory that a is (or whatever you are pointing to)
	a = 66
	fmt.Println(a, c)
	fmt.Println(&a, c) // Verify same memory location
	// Derefencing below
	fmt.Println(a, *c) // Derefence the pointer (drill down to the memory location and get whatever is stored there)
	*c = 14            // Will changa a and c
	fmt.Println(a, *c)

	// Notice below how there is a byte in between each index in memory
	d := [3]int{1, 2, 3}
	e := &d[0]
	f := &d[1]
	fmt.Printf("%v %p %p\n", d, e, f)

	pointersWithoutSettingType()

	variableHandling()
}

type myStruct struct {
	foo int
}

func pointersWithoutSettingType() {
	fmt.Println("Without setting type")
	var ms *myStruct
	ms = &myStruct{foo: 42} // object initialization
	// in example above we have a myStruct in ms that ALSO has the foo field.
	fmt.Println(ms)

	// You can also use new function (will instantiate it without data)
	var ms2 = new(myStruct) // In this example our ms2 does not contain the field foo. But as shown below - go understands that we want to talk to the original object and its field foo
	fmt.Println(ms2)
	/*
		Pointers that are initalized without data (before new or object initalization) have nil value (null)
	*/
	// Derefencing to get the value
	// (*ms2).foo = 42
	// fmt.Println((*ms2).foo)
	// Above would be logical but is not needed in go
	ms2.foo = 42 // this gives same result as above
	fmt.Println(ms2.foo)
}

func variableHandling() {
	a := [3]int{1, 2, 3}
	b := a
	fmt.Println(a, b)
	a[1] = 42
	fmt.Println(a, b)
	// as mention in array and slices above is a copy of the data
	// below (slice) is referencing the same data
	c := []int{1, 2, 3}
	d := c
	fmt.Println(c, d)
	c[1] = 42
	fmt.Println(c, d)

	// Maps also works as reference data (pointers)
}
