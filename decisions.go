package main

import "log"

func main() {
	myNum := 100
	myTrue := false
	myVar := "сat"

	// if/else
	if myNum > 99 && myTrue {
		log.Println("First Block")
	} else {
		log.Println("Second Block")
	}

	// switch
	switch myVar {
	case "сat":
		log.Println("this is cat")
	case "dog":
		log.Println("this is dig")
	default:
		log.Println("it nothing")
	}
}
