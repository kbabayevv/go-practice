package structs

import (
	"errors"
	"fmt"
	"time"
)

func (u *User) outPutDetails() {
	fmt.Println("User Details: "+u.firstName, u.lastName, u.birthDate)
}

func (u *User) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthDate string) (User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return User{}, errors.New("firstName, lastName and BirthDate are required")
	}
	return User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func Structs() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := newUser(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.outPutDetails()
	appUser.clearUserName()
	appUser.outPutDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	_, _ = fmt.Scanln(&value)
	return value
}

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}
