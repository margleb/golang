package config

import (
	"text/template"
)

// конфиг настройки приложения
type AppConfig struct {
	UseCache bool
	// кеш шаблонов
	TemplateCache map[string]*template.Template
}
