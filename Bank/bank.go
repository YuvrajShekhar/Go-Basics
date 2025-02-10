package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	for {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your Choice : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your Deposit :")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than zero.")
				continue
			}

			if depositAmount > accountBalance {
				fmt.Println("Invalid amount. You can withdraw more than you have.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Your new balance is", accountBalance)
		case 3:
			fmt.Print("Your Withdrawal amount :")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than zero.")
				return
			}
			accountBalance -= withdrawAmount
			fmt.Println("Your new balance is", accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for choosing our bank!!")
			return
		}
	}

}
