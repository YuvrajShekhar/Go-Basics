package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	user_firstName := getUserData("Please enter your first name: ")
	user_lastName := getUserData("Please enter your last name: ")
	user_birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	appUser, err := user.New(user_firstName, user_lastName, user_birthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "pass")

	admin.User.OutputUserDetails()
	admin.User.ClearUsername()
	admin.User.OutputUserDetails()

	// ... do something awesome with that gathered data!
	appUser.OutputUserDetails()
	appUser.ClearUsername()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
