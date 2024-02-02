package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	// `` states to tell go when recieve/creating jason in this format what to use here's
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	myJson := `
[
    {
        "first_name": "Clark",
        "last_name": "Kent",
        "hair_color": "black",
        "has_dog": true
    },
    {
        "first_name": "Bruce",
        "last_name": "Wayne",
        "hair_color": "black",
        "has_dog": false
    }
]`

	var unmarshalled []Person

	//Get data from json into slice/struct Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	// write json from a struct
	var mySlice []Person
	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColor = "black"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	//MarshalIndent for readability not for production
	newJson, err := json.MarshalIndent(mySlice, "", "     ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	//Print and cast to string
	fmt.Println(string(newJson))

} //End of main
