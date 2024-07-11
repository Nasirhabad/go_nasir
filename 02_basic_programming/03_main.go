package main

import (
	"fmt"
)

func reverseString(input string) string {
	// Convert string to rune slice to handle multi-byte characters
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var input string
	fmt.Println("Enter a string:")
	fmt.Scanln(&input)

	reversed := reverseString(input)
	fmt.Println("Reversed string:", reversed)
}
