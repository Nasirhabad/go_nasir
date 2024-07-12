package main

import (
	"errors"
	"fmt"
)

func main() {
	// function -> is a reusable unit
	// 1 -> remove code duplicates
	// 2 -> split the larger code base into many parts

	// function
	// 1 -> void (without returning any value)
	// 2 -> returns a value
	var a, b int

	fmt.Print("insert first number: ")
	fmt.Scan(&a)

	fmt.Print("insert second number: ")
	fmt.Scan(&b)

	result := sum(a, b)

	fmt.Println("the sum result: ", result)

	fmt.Println(multiSum(0, 1, 2))
	fmt.Println(multiSum(0, 1, 2, 3, 4, 5, 76, 8, 9, 2, 1, 1, 1, 1, 1, 1))

	res, err := div(12, 0)

	// nil -> empty value (reference type)

	// error-handling
	if err != nil {
		fmt.Println("an error occurred: ", err)
	} else {
		fmt.Println("the division result: ", res)
	}
}

func sum(a, b int) int {
	result := a + b
	return result
}

// multiSum returns a sum of a numbers
func multiSum(init int, nums ...int) int {
	sum := init

	for _, num := range nums {
		sum += num
	}

	return sum
}

// multiple returns (error handling)
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return a / b, nil
}
