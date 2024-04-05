package main

import "fmt"

/*
We might want to use type aliases to make our code more readable
Especially when we are working with complex data structures like maps
*/
type intMap map[string]int

/*
This allows us to define methods on the type
*/
func (m intMap) output() {
	fmt.Println(m)
}

func main() {
	/*
		Maps
		Maps are key-value pairs
		Keys can be of any type that supports equality
	*/
	websites := map[string]string{
		"Google":   "https://www.google.com",
		"Facebook": "https://www.facebook.com",
	}

	// Maps are always dynamic
	websites["Twitter"] = "https://www.twitter.com"
	websites["LinkedIn"] = "https://www.linkedin.com"

	fmt.Println(websites)

	// Extracting
	fmt.Println(websites["Google"])

	// Deleting
	delete(websites, "Facebook")
	fmt.Println(websites)

	// Make function
	marks := make(intMap, 4)
	marks["John"] = 90
	marks["Doe"] = 85
	marks["Jane"] = 95
	marks["Alice"] = 100
	marks.output()

	// Iterating
	for key, value := range marks {
		fmt.Println(key, value)
	}
	for key := range marks {
		fmt.Println(key)
	}
}
