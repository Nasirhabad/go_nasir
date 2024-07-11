package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	reversed := reverse(input)
	fmt.Println("Reversed string:", reversed)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}