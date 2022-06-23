package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/margleb/go-course/pkg/config"
	"github.com/margleb/go-course/pkg/handlers"
	"net/http"
)

// Routes - маршруты с мультиплексером
func routes(app *config.AppConfig) http.Handler {
	// маршрут
	mux := chi.NewRouter()
	// посредник - поглащает панику и печает стек
	mux.Use(middleware.Recoverer)
	// посредник - просто пишет в консоль
	// mux.Use(WriteToConsole)
	// посредник - устанавливаем CSRF токен
	mux.Use(NoSurf)

	// уст. маршруты
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	// возращаем
	return mux
}
