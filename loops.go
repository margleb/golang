package main

import "log"

type User struct {
	FirstName string
	LastName  string
	Email     string
}

func main() {

	mySlice := []string{
		"Cat", "Dog", "Kitten",
	}

	myMap := make(map[string]string)

	var myStruct []User

	s1 := User{
		FirstName: "Ivan",
		LastName:  "Urgant",
	}

	s2 := User{
		FirstName: "Gleb",
		LastName:  "Marianov",
	}

	myStruct = append(myStruct, s1)
	myStruct = append(myStruct, s2)

	myMap["name"] = "Gleb"
	myMap["lastName"] = "Martianov"
	myMap["age"] = "28"

	// 1. демонстрация простого for loop
	for i := 0; i < 10; i++ {
		log.Println(i)
	}

	// 2. перебор c slice (сортирован)
	for _, value := range mySlice {
		log.Println(value)
	}

	// 3. перебор с map (не сортирован)
	for i, value := range myMap {
		log.Println(i, ": ", value)
	}

	// 4. демонстрация со struct
	for _, user := range myStruct {
		log.Println("This is struct: ", user.FirstName)
	}

}
