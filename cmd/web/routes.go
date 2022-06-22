package main

import (
	"github.com/bmizerany/pat"
	"github.com/margleb/go-course/pkg/config"
	"github.com/margleb/go-course/pkg/handlers"
	"net/http"
)

// Routes - маршруты с мультиплексером
func routes(app *config.AppConfig) http.Handler {
	// мультиплексер
	m := pat.New()
	// маршруты
	m.Get("/", http.HandlerFunc(handlers.Repo.Home))
	m.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return m
}
