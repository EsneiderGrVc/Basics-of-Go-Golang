package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func threeArgs(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func returnTwoValues(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hello World")
	threeArgs(10, 100, "Bonyorno")
	value := returnValue(5)
	fmt.Println("Value: ", value)
	value1, value2 := returnTwoValues(10)
	fmt.Printf("First value is: %d\n Second value is: %d", value1, value2)
}
