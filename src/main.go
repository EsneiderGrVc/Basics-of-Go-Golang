package main

import "fmt"

func main() {
	// Constants Declarations

	const pi float64 = 3.64
	const pi2 = 3.6415

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	// Variables Declarations

	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero Values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Square Area
	const width = 10
	squareArea := width * width

	fmt.Println("Square Area: ", squareArea)
}
