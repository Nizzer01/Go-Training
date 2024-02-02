package helpers

import (
	"math/rand"
	"time"
)

/*//Create some types in this package
type SomeType struct {
	TypeName   string
	TypeNumber int
}

//Note: We can also defin functions here and call from this imported package
*/

//Create random number for chan
func RandomNumber(n int) int {
	//Need to seed to get an actual random number from our pool of 0-10
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}
