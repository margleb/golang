package main

import (
	"net/http"
)

// не может быть изменяемым
const portNumber = ":8080"

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(portNumber, nil)
}
