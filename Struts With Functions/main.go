package main

import "log"

type myStrut struct {
	FirstName string
}

//Func tied to struct via this receiver: (m *myStrut)
func (m *myStrut) printFirstName() string {
	//Can contain further business logic
	return m.FirstName
}

func main() {
	var myVar myStrut
	myVar.FirstName = "John"

	myVar2 := myStrut{
		FirstName: "Mary",
	}

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar is set to", myVar2.printFirstName())
}
