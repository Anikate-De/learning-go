// This file is also in the main package and is therefore accessible by bank.go without any import statement
package main

import "fmt"

func showOptions() {
	fmt.Println(
		`What would you like to do?
1. Check Balance
2. Deposit Money
3. Withdraw Money
4. Exit`)
}
