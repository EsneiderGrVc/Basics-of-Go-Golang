package main

import "fmt"

func main() {
	// For loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// For While loop
	fmt.Println("--------------------")
	var count int
	for count < 10 {
		fmt.Println(count)
		count++
	}

	// For Forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}
