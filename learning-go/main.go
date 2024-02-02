//package learning_go
package main

import (
	"log"
	"time"
)

//Define a type
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	//Vars & functions
	/*	fmt.Println("Hello, world.")

		var whatToSay string
		var i int

		whatToSay = "Goodbye, cruel world"
		fmt.Println(whatToSay)

		i = 7
		fmt.Println("i is set to: ", i)

		whatWasSaid, theOtherThingThatWasSaid := saySomething()

		fmt.Println("The function returned", whatWasSaid, theOtherThingThatWasSaid)*/

	//Types & Struts

	user := User{
		FirstName:   "Trev",
		LastName:    "Sawler",
		PhoneNumber: "1 555 555-2211",
	}

	log.Println(user.FirstName, user.LastName, "BirthDate:", user.BirthDate)
}
