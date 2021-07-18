package main

import "fmt"

func switchOddCheck(number int) {
	module := number % 2
	switch module {
	case 0:
		fmt.Printf("The number %d is Even", number)
	default:
		fmt.Printf("The number %d is Odd", number)
	}
}

func main() {
	switchOddCheck(7)
}
