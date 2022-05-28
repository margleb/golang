package main

import (
	"github.com/margleb/golang/helpers"
	"log"
)

func main() {
	log.Println("Hello")

	// импортируем из пакета helpers
	var myVar helpers.SomeType
	myVar.TypeName = "Some Name"
	log.Println(myVar.TypeName)
}
