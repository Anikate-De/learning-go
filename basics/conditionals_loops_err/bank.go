package main

import (
	"errors"
	"fmt"
	"os"      // FileSystem Read-Write Operations
	"strconv" // Provides string parsing functions
)

const balanceFileName = "balance.txt"

func saveFile(balance float64) {
	balText := fmt.Sprint(balance)
	// We cast the string into a collection of bytes
	// https://www.redhat.com/sysadmin/linux-file-permissions-explained
	os.WriteFile(balanceFileName, []byte(balText), 0644)
}

func loadBalance() (float64, error) {
	balBytes, err := os.ReadFile(balanceFileName)

	/*
		Error Handling in Go is much different from other languages that usually aim for catching exceptions
		In Go, errors are returned (usually as the last value) from functions

		All functions are therefore required to return a set of default values for all variables, regardless of errors thrown

		We can create new errors with messages, and check if errors have occurred or not using
		err != nil

		Hence, errors don't cause an app crash even if ignored unlike other languages
		However, to deliberately cause a crash, we use the panic function
		panic(err)
	*/
	if err != nil {
		return 1000, errors.New("unable to find balance data file, assuming new customer")
	}

	balText := string(balBytes)
	accountBalance, err := strconv.ParseFloat(balText, 64)

	if err != nil {
		return 1000, errors.New("unable to read balance from file, mismatched versions")
	}

	return accountBalance, nil
}

func main() {

	accountBalance, err := loadBalance()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}

	fmt.Println("Welcome to Go Bank!")

	/*
		Go only supports for-loops with similar syntax
		for var i = 0; i < 100; i++ {}

		If we want an infinite loop, we can declare a loop with just
		for {}

		For conditioned loops
		for someCondition {}
	*/
	for {

		fmt.Println(
			`What would you like to do?
1. Check Balance
2. Deposit Money
3. Withdraw Money
4. Exit`)

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
				saveFile(accountBalance)
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
				saveFile(accountBalance)
			}
		} else {
			break
		}
	}

	fmt.Println("Thank you for choosing our bank")
}
