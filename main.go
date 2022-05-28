package main

import (
	"github.com/margleb/golang/helpers"
	"log"
)

var rangeNum int = 10

func calculateValue(myChannel chan int) {
	// получаем любое рандомное число
	randNum := helpers.RandomNumber(rangeNum)
	// передаем рандомное число в канал
	myChannel <- randNum
}

func main() {
	// создаем канал
	intChannel := make(chan int)
	// закрываем канал после того как функция будет завершена
	defer close(intChannel)

	// передает канал в функ-цию
	go calculateValue(intChannel)

	// получаем значение из канала
	// (слушает канал и кидает в него значение)
	getNum := <-intChannel

	log.Println("Полученное число из канала:", getNum)
}
