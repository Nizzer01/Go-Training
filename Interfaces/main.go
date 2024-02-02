package main

import "fmt"

/*One function that accepts 2 different types
as they satisfy the interface requirements for type animal*/

//Define an interface
type Animal interface {
	//Add list of functions
	Says() string
	NumberOfLegs() int
}

//Define some structs
type Dog struct {
	Name  string
	Breed string
}
type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {

	dog := Dog{
		Name:  "Samson",
		Breed: "German Shepherd",
	}
	//Set as reference via &
	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "grey",
		NumberOfTeeth: 38,
	}
	//Set as reference via &
	PrintInfo(&gorilla)

}

//Define some functions

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

//Add receivers (note most recievers are of type pointer
func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "Ugh"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}
