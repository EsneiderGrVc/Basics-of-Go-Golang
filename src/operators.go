package main

import "fmt"

func main() {

	const width = 10
	squareArea := width * width

	fmt.Println("Square Area: ", squareArea)

	// Addition
	x := 10
	y := 50

	result := x + y
	fmt.Println("The sum is: ", result)

	// Substraction
	result = y - x
	fmt.Println("The substraction is: ", result)

	// Multiplication
	result = x * y
	fmt.Println("The multiplication is: ", result)

	// Divide
	result = y / x
	fmt.Println("Divide is: ", result)

	// Module
	result = y % x
	fmt.Println("Module is: ", result)

	// Increment
	x++
	fmt.Println("x++: ", x)

	// Decrement
	x--
	fmt.Println("x--: ", x)
}
