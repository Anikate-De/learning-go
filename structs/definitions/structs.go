package main

import (
	"fmt"

	"de.anikate/structs/user"
)

func main() {
	firstName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	birthDate := getUserData("Enter your birth date: (dd/mm/yyyy) ")

	// Create a new user of type user
	// If adding all fields, naming them is optional
	/* var newUser = user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	} */

	var newUser user.User = user.New(firstName, lastName, birthDate)

	/* outputUserDetails(&newUser) */

	// Call the method on the struct
	newUser.OutputUserDetails()

	newUser.ClearUserName()
	newUser.OutputUserDetails()

	// Create a new admin of type admin
	var admin user.Admin = user.NewAdmin(firstName, lastName, birthDate, "admin@localhost", "admin")
	admin.OutputUserDetails()
	admin.OutputAdminDetails()
}

func getUserData(prompt string) string {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)
	return input
}
