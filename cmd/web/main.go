package main

import (
	"github.com/margleb/go-course/pkg/handlers"
	"net/http"
)

// не может быть изменяемым
const portNumber = ":8080"

// main is the main application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_ = http.ListenAndServe(portNumber, nil)
}
