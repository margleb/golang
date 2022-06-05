package render

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

// массив функций применяемых к шаблону
var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	// получаем кеш страниц
	_, err := RenderTemplateTest(w)
	if err != nil {
		fmt.Println("Error getting template cache:", err)
	}

	parseTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err = parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {

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
		fmt.Println("The current page:", page)

		// парсим конкретную страницу
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// сканируем все страницы c *.layouts.tmp
		mathes, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		// если есть шаблоны, то парсим их
		if len(mathes) > 0 {
			// парсим к странице конкретный шаблон
			ts, err = ts.ParseGlob("/templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		// добавляем их в кеш
		myCache[name] = ts

	}

	return myCache, nil

}
