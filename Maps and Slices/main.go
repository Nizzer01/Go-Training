package main

import "log"

//Create a User struct Use with map examples
/*type User struct {
	FirstName string
	LastName  string
}*/

func main() {
	//Basic ways of storing data
	/*	var myString string
		var myInt int

		myString = "Hi"
		myInt = 11
		mySecondString := "another string"

		log.Println(myString, mySecondString, myInt)*/

	//Create a map of string
	/*	myMap := make(map[string]string)

		myMap["dog"] = "Samson"
		myMap["other-dog"] = "Cassie"

		//overwrite one of the key values
		myMap["dog"] = "fido"

		log.Println(myMap["dog"])
		log.Println(myMap["other-dog"])*/

	//Create a map of string int
	/*	myMap := make(map[string]int)

		myMap["First"] = 1
		myMap["Second"] = 2

		log.Println(myMap["First"])
		log.Println(myMap["Second"])*/

	//Maps can store anything eg types
	/*	//Maps are immutable
		//Maps are not sorted must be looked up by key
		//Syntax: map[what you want to look up by eg think key] datatype

		myMap := make(map[string]User)

		me := User{
			FirstName: "Trev",
			LastName:  "Sawler",
		}

		myMap["me"] = me

		log.Println(myMap["me"].FirstName)
		log.Println(myMap["me"].LastName)*/

	//Slices
	//Define a slice via []
	/*	var mySlice []int

		//Add data to slice
		mySlice = append(mySlice, 2)
		mySlice = append(mySlice, 1)
		mySlice = append(mySlice, 3)

		//Sort a slice
		sort.Ints(mySlice)

		log.Println(mySlice)*/

	//Using shorthand & ranges
	/*	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		log.Println(numbers)

		//Look at a range eg first 3
		//Note slice runs from 0....
		log.Println(numbers[0:2])

		//Mid range
		log.Println(numbers[6:9])*/

	//Slice of different types eg strings
	names := []string{"one", "seven", "fish", "cat"}
	log.Println(names)
}
