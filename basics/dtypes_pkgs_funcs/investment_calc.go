package main

// The recommended way to import multiple packages is to use the following syntax
import (
	"fmt"
	"math"
)

/*
Constants
Values that cannot change during the execution of the program, declared using `const`
We always need to assign an initial value

const constName type = value
const constName = value

Constants can be declared in a block:
const (

	constName1 = value1
	constName2 = value2

)
const (

	constName1 type = value1
	constName2 type = value2

)
*/
const inflationRate = 2.5

func main() {

	/*
		Basic Types
		int, float64, string (""|``), bool (true|false)

		Additional Types
		uint => An unsigned int (strictly non-negative integer)
		int32 => A 32-bit signed integer, from -2,147,483,648 to 2,147,483,647
		rune => An alias for int32, represents a Unicode code point (i.e., a single character), and is used when dealing with Unicode characters (e.g., 'a', 'ñ', '世')
		uint32 => A 32-bit unsigned integer, 0 to 4,294,967,295
		int64 => A 64-bit signed integer for larger range of values, from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
		int8 and uint8...

		Any Type
		any => The variable can hold a dynamic type

		Null Values
		All Go value types come with a "null value" by default if no value is set.

		int var would have a default value of 0 (because 0 is the null value of int, int32 etc):

		List of the null values for the different types:
		int => 0
		float64 => 0.0
		string => "" (an empty string)
		bool => false
	*/

	/*
		Type Inference and Casting
		Variables types are inferred from the value upon declaration
		We can use functions like int(), float64(), string() to cast variables

		We can also use the following syntax to declare a variable with a specific type:
		var varName type = value
	*/

	/*
		Variables
		Values that can change during the execution of the program, declared using `var`

		var varName type = value
		var varName = value
		var varName type
		varName := value (this is the recommend short-hand syntax for inferred types)

		Multiple variables can be declared in a single line:
		var varName1, varName2 type = value1, value2 (value1 and value2 must be of the same type)
		var varName1, varName2 = value1, value2 (each type is inferred)
	*/
	var investmentAmt float64
	var years float64 = 10
	expRetRate := 5.5

	/*
		To get std input, we use the `Scan()` function with a pointer to the variable
		Limitation: You can't (easily) fetch multi-word input values.
	*/
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmt)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expRetRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmt, expRetRate, years)

	/*
		The fmt package provides the function Printf() to format the output based upon the given formatter
		Or we could use comma-separated values in the print and println functions
	*/
	fmt.Println("Future Value:", futureValue)

	// Find the list of all formatters in the Golang Docs
	// https://pkg.go.dev/fmt
	fmt.Printf("Real Future Value (adjusted to inflation): %.2f\n", futureRealValue)

	// We can use Sprint or Sprintf functions to get the results of formatting
	formatFRVal := fmt.Sprintf("Real Future Value (adjusted to inflation): %.2f", futureRealValue)
	fmt.Println(formatFRVal)

	// Multiline Strings (using backticks)
	// The escape characters like \n are treated as plaintext, tabs are included
	multilineStr := `This is on a
	seperate line\n`
	fmt.Print(multilineStr)
}

/*
Functions need to be defined as

	func funcName(param1 type, param2 type) (returnType1, returnType2) {
		...
		return (val1, val2)
	}

You can chain parameters
func funcName(p1, p2, p3 type) {}
*/

/*
Alternative Return

	func ... (futureValue float64, futureRealValue float64) {
		futureValue = ... (we do not declare variables here, as the return type does that)
		futureRealValue = ...
		return (go automatically return the list of variables)
	}
*/
func calculateFutureValues(investmentAmt, expRetRate, years float64) (float64, float64) {
	futureValue := investmentAmt * math.Pow((1+expRetRate/100), years)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	return futureValue, futureRealValue
}
