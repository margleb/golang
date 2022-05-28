package main

import "log"

type myStruct struct {
	FullName string
}

func (u *myStruct) fullName() string {
	return u.FullName
}

func main() {

	// 1 ый вариант присвоения
	var user myStruct
	user.FullName = "Gleb Martianov"
	// log.Println("first example: ", user.FullName)
	log.Println("first example: ", user.fullName())

	// 2 ой вариант
	user2 := myStruct{
		FullName: "Иван",
	}
	// log.Println("second example:", user2.FullName)
	log.Println("first example: ", user2.fullName())

}
