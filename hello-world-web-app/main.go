package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

//Home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Home page")
}

//About page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the About page and 2 + 2 is: %d", sum))
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by zero")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))
}

func divideValues(x, y float32) (float32, error) {

	if y <= 0 {
		err := errors.New("cannot divide by 0")
		return 0, err
	}
	result := x / y
	return result, nil

}

//Adds 2 int's and returns sum
func addValues(x, y int) int {
	return x + y
}

//Main app functions
func main() {

	/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World!")

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprint("Number of bytes written: %d", n))
	})*/

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	//Start web server got to localhost:8080 in browser
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
