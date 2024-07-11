package main

import (
	"fmt"
)

func main() {
	// Example values
	a := 5.0 // length of the first parallel side
	b := 7.0 // length of the second parallel side
	h := 4.0 // height

	// Step 3: Process
	area := ((a + b) * h) / 2

	// Step 4: Output
	fmt.Printf("The area of the trapezoid with a = %.2f, b = %.2f, and h = %.2f is: %.2f\n", a, b, h, area)
}
