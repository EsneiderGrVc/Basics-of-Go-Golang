package main

import "fmt"

func main() {
	// Array
	var array [4]int

	// Assign index 0 and 1
	array[0] = 1
	array[1] = 2

	// Print array
	fmt.Println(array)

	// Print len and capacity
	fmt.Printf("Lenght of array is: %d\nCapacity of array is: %d\n", len(array), cap(array))

	fmt.Println("-------------------------")

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	//Slice Methods
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append Method
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append new list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
