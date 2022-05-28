package main

import "log"

func main() {

	var name string
	var age int

	name, _ = nameAndAge("Gleb")
	_, age = nameAndAge("Ivan")

	// 1. передача в функцию
	log.Println("My name is ", name)
	// 2. передача функции в качестве параметра
	// log.Println("No, my name is ", sayMyName("Иван"))
	// 2. вывод возраста
	log.Println("My age is ", age)

}

func nameAndAge(name string) (string, int) {
	// возращает два значения имя и возраст
	return name, 28
}
