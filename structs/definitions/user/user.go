package user

import (
	"fmt"
	"time"
)

// Define a struct type User
// The fields are private by default, unless they start with a capital letter
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time // This is a time.Time type, which itself is a struct
}

/*
Struct Embedding

You can embed a struct within another struct, which is similar to inheritance in OOP
If we embed the struct anonymously, we can access the fields directly
*/
type Admin struct {
	email    string
	password string
	User     // This is anonymous embedding
	// User User // This is normal embedding, notice how the field name also begins with a capital letter
}

/*
// Although it is trivial to pass a pointer of the struct to save memory, we have used it here
func outputUserDetails(newUser *user) {
	// Go has a shorthand syntax for dereferencing struct pointers
	// This is the same as (*newUser).firstName, etc.
	fmt.Println(newUser.firstName, newUser.lastName, "was born on", newUser.birthDate, "and was created on", newUser.createdAt)
}
*/

// Attaching methods to a struct
func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, "was born on", u.birthDate, "and was created on", u.createdAt)
}

func (a Admin) OutputAdminDetails() {
	fmt.Println(a.email, ":", a.password)
}

// Methods can also modify the struct
// It's mandatory to pass a pointer to the struct if you want to modify it
// This is because Go passes arguments by value, so the struct would be copied
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

/*
Constructor Function
This is not a feature in Go, but a general practice to create a function that returns a struct

Since this constructor is already inside a dedicated user package, it is not necessary to name it NewUser
*/
func New(firstName, lastName, birthDate string) User {
	/*
		In Go, returned values are also passed around as copies
		Hence, it would make sense to return a pointer instead
		Code would look like this:
		func newUser(...) *user {
			...
			return &user{...}}

		var newUser *user = newUser(firstName, lastName, birthDate)


		Constructors functions can also have validation logic
		You may return error types if the input is invalid
	*/
	return User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
}

func NewAdmin(firstName, lastName, birthDate, email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User:     New(firstName, lastName, birthDate),
	}
}
