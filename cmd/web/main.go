package main

import (
	"github.com/margleb/go-course/pkg/config"
	"github.com/margleb/go-course/pkg/handlers"
	"github.com/margleb/go-course/pkg/render"
	"log"
	"net/http"
)

// не может быть изменяемым
const portNumber = ":8080"

// main is the main application function
func main() {

	// настройки приложения
	var app config.AppConfig

	// получаем кеш шаблонов
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create tmp cache")
	}
	// сохраняем его в гл. переменную TemplateCache
	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_ = http.ListenAndServe(portNumber, nil)
}
