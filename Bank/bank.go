package main

import "fmt"

func main() {
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("1. Deposit Money")
	fmt.Println("1. Withdraw Money")
	fmt.Println("1. Exit")

	var choice int
	fmt.Print("Your Choice : ")
	fmt.Scan(&choice)

	fmt.Println("Your choice : ", choice)

}
