package main

import "fmt"

func main() {
	// Array of strings
	var productNames = [5]string{"Pen", "Pencil", "Marker", "Eraser", "Sharpener"}
	prices := [5]float64{10.5, 20.5, 30.5, 40.5, 50.5}
	fmt.Println(productNames)
	fmt.Println(prices)

	// Indexing
	fmt.Println(productNames[1], prices[1])

	// Modifying
	productNames[1] = "BallPen"
	prices[1] = 25.5
	fmt.Println(productNames[1], prices[1])

	// Slicing
	fmt.Println(productNames[2:4])
	fmt.Println(prices[:4])
	fmt.Println(productNames[2:])
	// !Note: Negative indexing is not supported in Go
	// !Note: The end index cannot be greater than the length of the array

	// If we modify the slice, the original array is also modified
	subProductNames := productNames[3:4]
	subProductNames[0] = "SketchPen"
	fmt.Println(productNames)

	// It can be said that a slice is a small window into the array, hence slicing also doesn't copy the array

	// Length: The number of elements in the slice/array
	fmt.Println(len(productNames))

	// Capacity: The number of elements in the underlying array of the slice, but only to the right of the slice
	fmt.Println(cap(productNames))

	/*
		Dynamic arrays
		In Go, we use slices to create dynamic arrays
		When we don't specify the length of the array, Go creates a slice
		And when the size of the slice exceeds the capacity of the underlying array, Go creates a new array with double the capacity and copies the elements to the new array
		And the old array is garbage collected
	*/
	names := []string{"John", "Doe", "Jane", "Doe"}
	fmt.Println(names)
	// names[4] = "Alice" // This will throw an error
	names = append(names, "Alice")
	fmt.Println(names)
	// Similarly, you can remove elements from the slice using slicing

	// Unpacking
	otherNames := []string{"Bob", "Eve"}
	names = append(names, otherNames...)
	fmt.Println(names)

	// Make function
	// Creates a slice with a length of 2 and a capacity of 5
	dogs := make([]string, 2, 5)
	dogs[0] = "Bulldog"
	dogs[1] = "Poodle"
	fmt.Println(dogs)

	// Iterating
	for index, value := range names {
		fmt.Println(index, value)
	}

	for index := range names {
		fmt.Println(index)
	}

}
