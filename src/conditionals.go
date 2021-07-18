package main

import "fmt"

func main() {
	value1 := 2
	value := 1
	if value1 == 1 {
		fmt.Println("It's 1")
	} else {
		fmt.Println("It isn't 1")
	}

	// Logic Operator: &&
	if value1 == 2 && value == 1 {
		fmt.Println("It's True")
	}

	// Logic Operator: ||
	if value1 == 0 || value == 1 {
		fmt.Println("One of those are True")
	} else {
		fmt.Println("No one is True")
	}
}
