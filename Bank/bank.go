package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your Choice : ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is", accountBalance)
	} else if choice == 2 {
		fmt.Print("Your Deposit :")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than zero.")
			return
		}

		if depositAmount > accountBalance {
			fmt.Println("Invalid amount. You can withdraw more than you have.")
			return
		}
		accountBalance += depositAmount
		fmt.Println("Your new balance is", accountBalance)
	} else if choice == 3 {
		fmt.Print("Your Deposit :")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)

		if withdrawAmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than zero.")
			return
		}
		accountBalance -= withdrawAmount
		fmt.Println("Your new balance is", accountBalance)
	} else {
		fmt.Println("Goodbye!")
	}

}
