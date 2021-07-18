package main

import (
	"fmt"
	"os"
)

func checkPass(argPass string) {

	var realpassword string = "helloworld"
	if argPass == realpassword {
		fmt.Println("Yeah! These is the password")
	} else {
		fmt.Println("Ops! Keep Trying")
	}
}

func main() {

	var password string = os.Args[1]
	checkPass(password)
}
