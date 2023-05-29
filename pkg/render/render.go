package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Dmkk01/bed-and-breakfast/pkg/config"
	"github.com/Dmkk01/bed-and-breakfast/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var ts map[string]*template.Template
	if app.UseCache {
		ts = app.TemplateCache
	} else {
		ts, _ = CreateTemplateCache()
	}

	t, ok := ts[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf, td)
	if err != nil {
		log.Println("Error executing template :", err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser :", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("../../templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}

// old logic
// var tc = make(map[string]*template.Template)

// func RenderTemplateTest(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	_, inMap := tc[t]

// 	if !inMap {
// 		log.Println("Template not in cache, parsing now...")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println("Error creating template cache :", err)
// 			return
// 		}
// 	} else {
// 		log.Println("Using template from cache")
// 	}

// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)

// 	if err != nil {
// 		fmt.Println("Error executing template :", err)
// 		return
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		"../../templates/" + t,
// 		"../../templates/base.layout.tmpl",
// 	}

// 	tmpl, err := template.ParseFiles(templates...)

// 	if err != nil {
// 		return err
// 	}

// 	tc[t] = tmpl

// 	return nil
// }
