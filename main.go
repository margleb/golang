package main

import (
	"errors"
	"fmt"
	"net/http"
)

// не может быть изменяемым
const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

// addValue adds to integers and return the sum
// со строчной буквы делает видимость фун-ции только в этом пакете
func addValues(x, y int) int {
	return x + y
}

// Devide - страница деления
func Divide(w http.ResponseWriter, r *http.Request) {

	f, err := divideValues(100.0, 0.0)

	if err != nil {
		fmt.Fprintf(w, "Cannot devide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f devided by %f is %f", 100.00, 0.0, f))
}

func divideValues(x, y float32) (float32, error) {

	if y <= 0 {
		err := errors.New("cannot devide by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/devide", Divide)

	// пишет в консоль
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
