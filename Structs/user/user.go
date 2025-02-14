package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User     User
}

func (u *User) ClearUsername() {
	u.firstName = ""
	u.lastName = ""
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First Name, Last name or birthdate not found")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
	}, nil
}
