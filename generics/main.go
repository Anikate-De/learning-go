package main

import "fmt"

func main() {
	fmt.Println(add(10, 20))
	fmt.Println(add(10.54, 20.245))
	fmt.Println(add(float32(10.524), float32(20.25)))
}

/*
Here, we define a generic function that accepts values of type `int`, `float32`, and `float64`.
And returns the sum of the two values, of the same type.

Multiple types can be defined using the `|` operator.

Multiple placeholders can be defined using the `,` operator.
func add[T, K] (a T, b K) T {
	return a + b
}


Thanks to generics in this example, we can be sure that the operands passed to the function can be added together.
*/
func add[T int | float32 | float64](a, b T) T {
	return a + b
}
