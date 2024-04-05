package main

import (
	"fmt"
)

func main() {
	printSomething(10)
	printSomething("Hello")
	printSomething(10.5)

	switchTypePrint(10)
	switchTypePrint("Hello")
	switchTypePrint(10.5)
	switchTypePrint(true)

	isInteger(10)
	isInteger("Hello")
}

/*
In this dummy function, we are accepting an `interface{}` as an argument.
Alternatively, we can use the `any` keyword

This allows us to pass any type of value to this function.
*/
func printSomething(value interface{}) {
	fmt.Println(value)
}

/*
In a switch-case statement, we can use the `value.(type)` syntax to check the type of the value passed to the function.
*/
func switchTypePrint(value any) {
	switch value.(type) {
	case int:
		fmt.Println("This is an integer")
	case string:
		fmt.Println("This is a string")
	case float64:
		fmt.Println("This is a float")
	default:
		fmt.Println("This is an unknown type")
	}
}

/*
value.(type) can not be used outside of a switch-case statement.
To check the type of a variable, we can use the type assertion syntax.
*/
func isInteger(value any) {
	typedVal, ok := value.(int) // typedVal will either be the value of the variable or the nil value of the type

	if ok {
		fmt.Println("This is an integer")
		fmt.Println("Value is:", typedVal)
	} else {
		fmt.Println("This is not an integer")
		fmt.Println("Tried Parsing, value is:", typedVal)
	}
}
