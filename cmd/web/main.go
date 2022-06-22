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
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// создаем сервер
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	// запускаем его
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
