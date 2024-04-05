/*
To execute: go run app.go
To build: go build app.go
Before building, we should initialize the module: go mod init url/hello_world
This will create a go.mod file in the directory with the name and go version.

A package declaration is absolutely necessary for every Go file.
Packages can have multiple files. And a project (module) can have multiple packages.

The main package is the entry point of the program.
If we build with a different package, it will not generate the build output

On Windows, the build output will be an .exe file.
On Linux and macOS, the build output will be an executable file without any extension.
*/
package main

import "fmt" // built-in package for formatted I/O

/*
The main function is the entry point of the program.
Any two programs in the main package cannot both have a main function.
`main is redeclared in this block` error will be thrown
*/
func main() {
	fmt.Print("Hello, World!") // Strings can only have either double quotes or backticks.
}
