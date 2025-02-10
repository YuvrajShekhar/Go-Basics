package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	balancebyte, err := os.ReadFile(accountFile)

	if err != nil {
		panic("Can't Continue sorry, failed to fetch balance file")
	}

	balanceString := string(balancebyte)
	balanceFloat, err := strconv.ParseFloat(balanceString, 64)

	if err != nil {
		return 1000, errors.New("failed to convert balance to float")
	}

	return balanceFloat, nil
}

func writeBalanceToFile(balance float64) {
	balance_text := fmt.Sprint(balance)
	os.WriteFile(accountFile, []byte(balance_text), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Print("ERORRRR")
		fmt.Println(err)
		fmt.Println("----------")
	}

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

			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
			fmt.Println("Your new balance is", accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for choosing our bank!!")
			return
		}
	}

}
