package main

import "fmt"

func oddNumbers(num int) {

	var count int

	for i := 0; i < num; i++ {
		if i%2 == 0 {
			count++
		}
	}
	fmt.Printf("The range of number %d has %d odd values", num, count)
}

func main() {
	oddNumbers(10)
}
