package main

import "log"

func main() {
	//Pointers overview
	// & for reference to a var
	// * create a pointer
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)

	//use ref in memory
	changeUsingPointer(&myString)
	log.Println("after func call myString is set to", myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
