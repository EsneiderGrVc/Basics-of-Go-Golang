package main

import "fmt"

func main() {

	// Variables and Declarations
	helloMessage := "Hello"
	worldMessage := "world"

	//fmt Println
	fmt.Println(helloMessage, worldMessage)

	//Printf
	name := "Platzi"
	courses := 500

	fmt.Printf("%s tiene más de %d cursos\n", name, courses)
	fmt.Printf("%v tiene más de %v cursos\n", name, courses)

	// Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", name, courses)
	fmt.Println(message)

	// Data Type
	fmt.Printf("helloMessage type: %T", helloMessage)

}
