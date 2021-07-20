package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) bool {
	text = strings.ToLower(text)

	var textReverse string
	var textLen int = len(text) - 1

	for i := textLen; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		return true
	} else {
		return false
	}
}

func main() {
	slice := []string{"hola", "que", "hace"}
	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	/* Following string within "isPalindrome" does'nt work
	thus has whitespaces */
	fmt.Println(isPalindrome("Was it a car or a cat I saw"))
}
