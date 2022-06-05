package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// массив функций применяемых к шаблону
var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	// получаем кеш страниц
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// берем шаблон из сгенерированного кеша
	t, ok := tc[tmpl]
	// если шаблон не найден
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	// записываем в буфер полученный шаблон
	_ = t.Execute(buf, nil)
	// записываем его в ответ
	_, err = buf.WriteTo(w)
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
