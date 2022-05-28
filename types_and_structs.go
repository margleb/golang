package main

import (
	"log"
	"time"
)

var s string = "outside"

type thisUser struct {
	UserName string
	Email    string
	BirthDay time.Time
}

func main() {

	var s1 string = "inside"
	noType := "without type"
	user := thisUser{
		UserName: "Gleb Martianov",
		Email:    "margleb.dm@yandex.ru",
	}

	// доступна в любой функции
	log.Println("outside var:", s)
	// доступна, в пределах только этой фун-ции
	log.Println("inside var:", s1)
	// переписанная фун-ция
	log.Println("overwrite var:", someFunction("xxx"))
	// без указанного типа переменная
	log.Println("without type var:", someFunction(noType))
	// полученное значение из struct
	log.Println("user name is: ", user.UserName)
	// получаем значение которого нет
	log.Println("user name is: ", user.BirthDay)

}

func someFunction(s string) string {
	return s
}
