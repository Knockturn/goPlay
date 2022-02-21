package main

// Course part start: https://www.youtube.com/watch?v=YS4e4q9oBaU&t=3425s
/*
AGENDA:
Boolean Types
Numeric Types (Integers, Floating point, Complex numbers)
Text Types
*/

import (
	"fmt"
)

func main() {
	a := true
	fmt.Printf("%v, %T\n", a, a)

	// Bool set by logical test
	fmt.Println("Bool Logival Tests:")
	b := 1 == 1
	c := 1 == 2
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)

	//Numeric Types
	fmt.Println("\nNumeric Types:")
	d := 10 // Binary: 1010
	fmt.Printf("%v, %T\n", d, d)
	var e uint = 3 // Binary 0011
	fmt.Printf("%v, %T\n", e, e)
	/*
	int8: -128 - 127
	int16: -32768 - 32767
	int32: -2 billion - 2 billion
	int64: forget about it...

	uint8: 0 - 255 (same as a byte)
	uint16: 0 - 65535
	uint32: 0 - 4294967295
	*/

	// arithmetic operations
	fmt.Println("\nArithmetic operations")
	fmt.Println(uint(d) + e)
	fmt.Println(uint(d) - e)
	fmt.Println(uint(d) / e) // Gives back an int (so it drops the remainder)
	fmt.Println(uint(d) % e) // Gives the remainder

	// Bit operator
	fmt.Println("\nBit operations")
	fmt.Println(uint(d) & e) // 0010 (&)
	fmt.Println(uint(d) | e) // 1011 (or)
	fmt.Println(uint(d) ^ e) // 1001 (Exclusive or)
	fmt.Println(uint(d) &^ e) // 0100 (and not)

	// Bit shifting
	f := 8 // 2^3
	fmt.Println("\nBit shifting")
	fmt.Println(f << 3) // 2^3 * 2^3 or 2^6
	fmt.Println(f >> 3) // 2^3 / 2^3 or 2^0

	// Floating point
	fmt.Println("\nFloating points nr.")
	g := 3.14
	g = 13.7e72 // Too high for float32 - so if you initialize it as a 32 manually it would fail with this line. Now it starts as a float64
	g = 2.1E14
	fmt.Printf("%v, %T\n", g, g)
	// NOTE: for floats Bit shifting, Bit operations does not work. Also remainder arithmatics does not work either. ALl those are only w. ints.

	// Complex number (float plus float)
	fmt.Println("\nComplex numbers")
	var h complex64 = 1 + 2i // i = imaginary number: https://golangdocs.com/complex-numbers-in-golang
	fmt.Printf("%v, %T\n", h, h)
	fmt.Printf("%v, %T\n", real(h), real(h))
	fmt.Printf("%v, %T\n", imag(h), imag(h))
	i := 2 + 2i
	j := 3 + 2i
	fmt.Println(i + j)
	fmt.Println(i - j)
	fmt.Println(i * j)
	fmt.Println(i / j)
	// Making a complex nr:
	var k complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", k, k)

	// Text Types
	// String = any UTF8 character - in go a string is an alias for a byte
	fmt.Println("\nText types:")
	l := "this is a string"
	fmt.Printf("%v, %T\n", l, l)
	fmt.Printf("%v, %T\n", l[2], l[2]) // Can be treated like an array
	m := "this is also a string"
	fmt.Printf("%v, %T\n", l + m, l + m)
	//String can be converted to "slice of bytes"
	/*
	NOTES:
	This is mostly used for when sending strings between computers or such. THen sending it as bytes is best.
	*/
	fmt.Println("Slice of bytes (collection of bytes):")
	n := []byte(l)
	fmt.Printf("%v, %T\n", n, n)

	// A Rune: Any UTF32 character - (UTF8 is also a valid UTF32). Declared by single quotes '
	/*
	Runes are type alias for int32. So below will return type int32 which might look weird
	A rune IS an int32.
	Some string documentation on reading runes: https://pkg.go.dev/strings#Reader.ReadRune
	*/
	fmt.Println("RUNES:")
	o := 'a' // It is not just an int32 because of the implecit declaration.
	fmt.Printf("%v, %T\n", o, o)
	var p rune = 'a'
	fmt.Printf("%v, %T\n", p, p)
}