package main

import "log"

func main() {

	//If statements basics
	isTure := false
	if isTure == true {
		log.Println("isTrue is", isTure)
	} else {
		log.Println("isTrue is", isTure)
	}

	cat := "cat"
	if cat == "cat" {
		log.Println("cat is", cat)
	} else {
		log.Println("cat is", cat)
	}

	myNum := 100
	isTure = false
	if myNum > 99 && !isTure {
		log.Println("myNum is greater than 99 and isTrue is set to true")
	} else if myNum < 100 && isTure {
		log.Println("1")

	} else if myNum == 101 || isTure {
		log.Println("2")
	} else if myNum > 100 && isTure == false {
		log.Println("3")
	}

	//Switch statements basics
	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")
	case "fish":
		log.Println("cat is set to fish")
	default:
		log.Println("cat is something else")
	}

}
