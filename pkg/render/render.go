package render

import (
	"bytes"
	"fmt"
	"github.com/margleb/go-course/pkg/config"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// массив функций применяемых к шаблону
var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	var tc map[string]*template.Template
	//  если необходимо использовать кеш
	if app.UseCache {
		// получаем значение кеша из ранее сгенерированного
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// берем шаблон из сгенерированного кеша
	t, ok := tc[tmpl]
	// если шаблон не найден
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	// записываем в буфер полученный шаблон
	_ = t.Execute(buf, nil)
	// записываем его в ответ
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

// CreateTemplateCache создаем кеш шаблона как map
func CreateTemplateCache() (map[string]*template.Template, error) {

	// создаем map кеш страниц
	myCache := map[string]*template.Template{}

	// сканируем все страницы *.page.tmp
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// перебираем кажду страницу
	for _, page := range pages {

		// название страницы
		name := filepath.Base(page)

		// парсим конкретную страницу
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// сканируем все страницы c *.layouts.tmp
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		// если есть шаблоны, то парсим их
		if len(matches) > 0 {
			// парсим к странице конкретный шаблон
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		// добавляем их в кеш
		myCache[name] = ts

	}

	return myCache, nil

}
