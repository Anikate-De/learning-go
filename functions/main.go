package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	transformNumbers(&numbers, square)
	fmt.Println(numbers)
	transformNumbers(&numbers, doubleFn())
	fmt.Println(numbers)

	// Anonymous functions
	transformNumbers(&numbers, func(n int) int {
		return n * 3
	})
	fmt.Println(numbers)

	sum := sum(1, 24, 241, 241, 41)
	fmt.Println(sum)
}

// Functions as parameters
func transformNumbers(numbers *[]int, operation func(int) int) {
	for key, val := range *numbers {
		(*numbers)[key] = operation(val)
	}
}

func square(n int) int {
	return n * n
}

func double(n int) int {
	return n * 2
}

// Functions as return values
func doubleFn() func(int) int {
	return double
}

// Variadic functions: Functions that can take a variable number of arguments
func sum(numbers ...int) int {
	result := 0
	for _, val := range numbers {
		result += val
	}
	return result
}
