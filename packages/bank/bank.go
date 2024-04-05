package main

import (
	"fmt"

	"de.anikate/bank/file_operations"    // To import packages, we need to give full path wrt to the module root
	"github.com/Pallinder/go-randomdata" // import 3rd party packages using go get
)

const balanceFileName = "balance.txt"

func main() {

	accountBalance, err := file_operations.ReadFloatFromFile(balanceFileName, 1000)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}

	fmt.Println("Welcome to Go Bank!\nReach out to us 24x7 at", randomdata.PhoneNumber())

	/*
		Go only supports for-loops with similar syntax
		for var i = 0; i < 100; i++ {}

		If we want an infinite loop, we can declare a loop with just
		for {}

		For conditioned loops
		for someCondition {}
	*/
	for {
		showOptions()

		var choice int

		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)

		/*
			AND: &&
			OR: ||

			Alternative: The switch-case-default statement

			switch varName {
			case val1:
			case val2:
			...
			default:
			}

			Note that unlike other languages (Java), only one case statement can be executed by a switch
			Also, break has its own application within switch scope and cannot be used to break out of parent scopes (if any)
		*/
		if choice == 1 {
			fmt.Printf("Your Balance is: %.2f\n", accountBalance)
		} else if choice == 2 {
			var deposit float64
			fmt.Print("Deposit Amount: ")
			fmt.Scan(&deposit)

			if deposit <= 0 {
				fmt.Println("Invalid Deposit Amount: Please enter a positive value.")
				continue
			} else {
				accountBalance += deposit
				file_operations.SaveFloatToFile(accountBalance, balanceFileName)
			}
		} else if choice == 3 {
			var withdraw float64
			fmt.Print("Withdraw Amount: ")
			fmt.Scan(&withdraw)

			if withdraw <= 0 {
				fmt.Println("Invalid Withdraw Amount: Please enter a positive value.")
				continue
			} else if withdraw > accountBalance {
				fmt.Println("Invalid Withdraw Amount: Insufficient Balance.")
				continue
			} else {
				accountBalance -= withdraw
				file_operations.SaveFloatToFile(accountBalance, balanceFileName)
			}
		} else {
			break
		}
	}

	fmt.Println("Thank you for choosing our bank")
}
