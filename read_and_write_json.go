package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   string `json:"age"`
}

func main() {

	myJson := `
[
	{
		"name": "Gleb",
		"email": "margleb.dm@yandex.ru",
		"age": "28"
	},
	{
		"firsName": "John",
		"email": "johnTomas.dm@gmail.com",
		"age": "30"
	}
]`

	// read json from a struct

	var unmarshalled []User

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshal json", err)
	}

	// log.Printf("unmarhseld %v", unmarshalled)

	// write json from a struct
	var mySlice []User

	var u1 User
	u1.Name = "Sergey"
	u1.Email = "sergeyIvanov@yandex.ru"
	u1.Age = "32"

	mySlice = append(mySlice, u1)

	var u2 User
	u2.Name = "Anton"
	u2.Email = "AnotonFedotov@yandex.ru"
	u2.Age = "25"

	mySlice = append(mySlice, u2)

	newJson, err := json.MarshalIndent(mySlice, "", "   ")
	if err != nil {
		log.Println("error marshal", err)
	}
	fmt.Println(string(newJson))

}
