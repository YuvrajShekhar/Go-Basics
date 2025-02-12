package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	presentOptions()
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Print("ERORRRR")
		fmt.Println(err)
		fmt.Println("----------")
	}

	for {

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

			accountBalance += depositAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Your new balance is", accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for choosing our bank!!")
			return
		}
	}

}
