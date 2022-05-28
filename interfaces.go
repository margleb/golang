package main

import "log"

// интерфейс
type Animal interface {
	makeSound() string // издает звук
	numberOfPaws() int // кол-во лап
}

// собака
type Dog struct {
	name  string
	breed string
}

// слон
type Elephant struct {
	hasTrunk    bool // есть ли хобот
	tailsNumber int  // кол-во хвостов
}

func main() {

	dog := Dog{
		name:  "Scooby",
		breed: "shepherd",
	}

	elephant := Elephant{
		hasTrunk:    true,
		tailsNumber: 4,
	}

	// проверяет наличие методов для struct
	printInfo(dog)
	printInfo(elephant)

}

func (d Dog) makeSound() string {
	return "woof"
}

func (d Dog) numberOfPaws() int {
	return 4
}

func (f Elephant) makeSound() string {
	return "weeeeeee"
}

func (f Elephant) numberOfPaws() int {
	return 5
}

func printInfo(a Animal) {

	log.Println("This animal make sound", a.makeSound(), "and has", a.numberOfPaws(), "paws")
}
