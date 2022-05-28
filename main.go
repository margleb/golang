package main

import (
	"fmt"
	"net/http"
)

func main() {

	// маршрут, по какому пути нужно ждать ответ
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// записываем ответ ответ в барузер
		n, err := fmt.Fprintf(w, "hello, world!")
		if err != nil {
			fmt.Println(err)
		}
		// Sprintf - позволят преобразовать в другой тип (int в string)
		// n - кол-во записанных байтов
		fmt.Println(fmt.Sprintf("Number of bytes written %d", n))
	})

	// слушаем конкретный порт на компьютере
	// если есть ошибка то не получаем ее
	_ = http.ListenAndServe(":8080", nil)

}
