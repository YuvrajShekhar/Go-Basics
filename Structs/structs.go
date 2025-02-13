package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (u *user) clearUsername() {
	u.firstName = ""
	u.lastName = ""
}

func (u user) outputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func newUser(firstName, lastName, birthdate string) (*user, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First Name, Last name or birthdate not found")
	}
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
	}, nil
}

func main() {
	user_firstName := getUserData("Please enter your first name: ")
	user_lastName := getUserData("Please enter your last name: ")
	user_birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user
	appUser, err := newUser(user_firstName, user_lastName, user_birthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ... do something awesome with that gathered data!
	appUser.outputUserDetails()
	appUser.clearUsername()
	appUser.outputUserDetails()
}

func outputUserDetails(u *user) {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
