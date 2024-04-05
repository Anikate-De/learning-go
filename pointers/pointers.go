package main

/*
Pointers
A variable that stores the memory address of another variable
Advantages:
- Prevent value copies
	When you pass a variable to a function, Go creates a copy of the variable
	This copy persists throughout the function's scope and is later removed by the Go Garbage Collector

	We can use pointers to pass the memory address of the variable instead of the value itself
	By doing this, we can modify the original variable's value

- Direct memory access
	We can directly access the memory address of a variable and modify its value
	Which can lead to less code (might be difficult to read however)


Go uses the same syntax for pointers as C and C++
var ptr *int
ptr = &someVar
fmt.Println(*ptr) // prints the value of someVar (dereferencing the pointer)
*/

import "fmt"

func main() {
	age := 40
	var agePtr *int = &age // Declare an age pointer

	fmt.Println("Age:", *agePtr)

	// All the memory addresses will be the same
	fmt.Println("Memory Address of Age:", &age)
	fmt.Println("Age Pointer:", agePtr)

	convertAgeToAdultAge(agePtr)
	fmt.Println("Memory Address of Adult Age:", &age)
	fmt.Println("Adult Age:", age)
}

// Receive a pointer to an integer
func convertAgeToAdultAge(age *int) {
	// We dereference the pointer to modify the value of the original variable
	*age = *age - 18
}
