package main

import (
	"log"
	"sort"
)

type myUser struct {
	Username string
	Email    string
}

func main() {
	// 1. созданный map
	// Maps НЕ могут быть сортированными
	myMapSting := make(map[string]string)
	// не рекомендуется
	// withInterface := make(map[string]interface{})
	myMapSting["first"] = "Gleb"
	myMapSting["second"] = "Ivan"
	// переписанное значение
	myMapSting["second"] = "Sergey"

	log.Println("First key sting: ", myMapSting["first"])
	log.Println("Second key sting: ", myMapSting["second"])

	// 2. int значения
	myMapInt := make(map[string]int)
	myMapInt["first"] = 1
	myMapInt["second"] = 2
	log.Println("First key int: ", myMapInt["first"])
	log.Println("Second key int: ", myMapInt["second"])

	// 3. добавляем struct
	myMapStruct := make(map[string]myUser)
	myUser := myUser{
		Username: "Ivan",
		Email:    "margleb.dm@yandex.ru",
	}
	myMapStruct["getUser"] = myUser
	log.Println("user struct is ", myMapStruct["getUser"].Email)

	// 4. c float значениями
	var floatNumber float32
	floatNumber = 1.1
	floatMap := make(map[string]float32)
	floatMap["floatNumber"] = floatNumber
	log.Println("Float number map: ", floatMap["floatNumber"])

	// Slices
	var mySlice []int
	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 6)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 11)
	// могут сортироваться
	sort.Ints(mySlice)

	mySliceString := []string{
		"Ivan",
		"Sergey",
		"Petya",
	}

	// диапазон
	log.Println(mySlice[2:4])
	log.Println(mySliceString)

}
