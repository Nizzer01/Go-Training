/*Overview of packages/Go modules*/
package main

import (
	"github.com/Nizzer01/myniceprogram/helpers"
	"log"
)

//Defining our own package/module
/*In term "go mod init github.com/Nizzer01/myniceprogram"

➜  Packages git:(main) ✗ go mod init github.com/Nizzer01/myniceprogram
go: creating new go.mod: module github.com/Nizzer01/myniceprogram

Ensure package for go inheritance is enabled in preferences*/

//func main() {

/*	//Basic using a package
	log.Println("Hello")

	//Use our type from our package
	var myVar helpers.SomeType
	myVar.TypeName = "Some name"
	log.Println(myVar.TypeName)*/

//PrintText("Hi ")
//}

/*func PrintText(s string) {
	log.Println(s)
}*/

//Using Channels to pass data between helpers and other packages
const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	//Pass to chnl via <-
	intChan <- randomNumber
}
func main() {
	//Define a chan
	intChan := make(chan int)

	//Close chn good practise (main will auto close)
	defer close(intChan)

	//Send as concurrent call
	go CalculateValue(intChan)

	//Listen for response from chnl
	num := <-intChan
	log.Println(num)

}
