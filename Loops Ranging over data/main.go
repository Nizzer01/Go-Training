package main

import "log"

func main() {

	//Basic loop overview

	//For loop
	for i := 0; i <= 5; i++ {
		log.Println(i)
	}

	//loop/range using range
	animals := []string{"dog", "fish", "horse", "cow", "cat"}
	for i, animal := range animals {
		log.Println(i, animal)
	}

	//loop using range & _ blank identifier
	for _, animal := range animals {
		log.Println(animal)
	}

	//loop/range using map
	animalsMap := make(map[string]string)
	//populate map
	animalsMap["dog"] = "Fido"
	animalsMap["cat"] = "Fluffy"

	for animalType, animal := range animalsMap {
		log.Println(animalType, animal)
	}

	//loop/range over strings
	//A sting is a slice of bytes output eg  0 : 79
	var firstLine = "Once upon a midnight dreary"

	for i, l := range firstLine {
		log.Println(i, ":", l)
	}

	//loop/range over custom objects

	//Create custom type
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smyth", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 20})
	users = append(users, User{"Sally", "Brown", "sally@brown.com", 45})
	users = append(users, User{"Alex", "Anderson", "alex@anderson.com", 17})

	//range over our custom user type
	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}

}
