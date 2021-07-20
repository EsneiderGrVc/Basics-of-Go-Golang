package main

import "fmt"

// Struct Definition
type car struct {
	brand string
	year  int
}

func main() {

	// car instance: First manner
	myCar := car{brand: "Toyota", year: 2021}
	fmt.Println(myCar)

	// car instance: Second instance
	var otherCar car
	otherCar.brand = "Ferrari"
	otherCar.year = 2020
	fmt.Println(otherCar)
}
