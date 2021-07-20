package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Jose"] = 14
	m["Pepito"] = 20

	// Walking a map
	for key, value := range m {
		fmt.Println(key, value)
	}

	// Print a specific value from a map
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
