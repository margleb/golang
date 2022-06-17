package config

import "text/template"

// конфиг. настройки приложения
type AppConfig struct {
	// кеш шаблонов
	TemplateCache map[string]*template.Template
}
