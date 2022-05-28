package main

import "log"

func main() {

	var name string
	name = "Gleb"
	// до применения поинтера
	log.Println("Before pointer ", name)
	// применяем pointer
	pointerChange(&name)
	log.Println("After pointer", name)

}

func pointerChange(s *string) {
	// адрес в памяти
	log.Println("адрес в памяти: ", s)
	// создаем новую переменную и меняем ее имя
	newName := "Sergey"
	*s = newName
}
