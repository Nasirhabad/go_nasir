package main

import "fmt"

func main() {
	// array -> store many items in specific size (fixed size)

	// approach 1
	var numbers [5]int

	// insert items
	numbers[0] = 12
	numbers[1] = 14
	numbers[2] = 46
	numbers[3] = 50

	// approach 1 for accessing array -> regular for loop
	// more customizable (specify which index)
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// approach 2 for accessing array -> for range
	for idx, val := range numbers {
		fmt.Println("the index: ", idx, " the value: ", val)
	}

}
