package main

// Course part start: https://www.youtube.com/watch?v=YS4e4q9oBaU&t=6473s

/*
AGENDA:
Arrays (Basis for Slices)
	Creation
	Built-in functions
	Working with arrays
Slices
	Creation
	Built-in functions
	Working with slices
*/

import (
	"fmt"
)

func main() {
	//Array: Has a fixed size that HAS to be known at compile time.
	grades := [4]int{97, 85, 93, 47}
	fmt.Printf("Grades: %v\n", grades)

	grades2 := [...]int{97, 85, 93, 47}
	fmt.Printf("Grades: %v\n", grades2)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Peter"
	students[2] = "Thomas"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #1: %v\n", students[1])
	fmt.Printf("Students size/length: %v\n", len(students))

	// Identity Matrix
	var identityMatrix [3][3]int = [3][3]int{[3]int{1,0,0}, [3]int{0,1,0}, [3]int{0,0,1}}
	fmt.Println(identityMatrix)

	//Copy an aray actually creates a new copy of it - and is NOT the same data
	a := [...]int{1,2,3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	// However - the & in front will mean that it points to the same data:
	c := &a
	c[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	//Slice: extends upon an array (does not have a fixed size at compile time). Also they are reference data so does not require the & to point to the same data.
	fmt.Println("Slices:")
	slice := []int{1,2,3}
	fmt.Printf("Slice: %v\n", slice)
	slice2 := slice
	slice2[1] = 21
	fmt.Println(slice)
	fmt.Println(slice2)
	/*
	NOTE: So be careful with slices. It will affect the underlying data!
	*/
	//Different ways of making a slice
	d := []int{1,2,3,4,5,6,7,8,9,10}
	e := d[:]	// slice of all elemnts
	f := d[3:]  // slice from 4th element to end
	g := d[:6]  // slice first 6 elements
	h := d[3:6] // slice the 4th, 5th and 6th elements - from and TO (but not including)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	f[0] = 666
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	//Last way to make a slice
	i := make([]int, 3)
	fmt.Println(i)
	fmt.Printf("Length: %v\n", len(i))
	fmt.Printf("Capacity: %v\n", cap(i))
	// Set max capacity
	j := make([]int, 3, 100)
	fmt.Println(i)
	fmt.Printf("Length: %v\n", len(j))
	fmt.Printf("Capacity: %v\n", cap(j))

	k := []int{}
	fmt.Println(k)
	fmt.Printf("Length: %v\n", len(k))
	fmt.Printf("Capacity: %v\n", cap(k))
	k = append(k, 1)
	fmt.Println(k)
	fmt.Printf("Length: %v\n", len(k))
	fmt.Printf("Capacity: %v\n", cap(k))

	l := make([]int, 3, 3)

	fmt.Printf("L Length: %v\n", len(l))
	fmt.Printf("L Capacity: %v\n", cap(l))

}

// Progress: https://youtu.be/YS4e4q9oBaU?t=7650