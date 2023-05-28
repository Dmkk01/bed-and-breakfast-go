package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("../../templates/"+tmpl, "../../templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Error executing template :", tmpl)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplateTest(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, inMap := tc[t]

	if !inMap {
		log.Println("Template not in cache, parsing now...")
		err = createTemplateCache(t)
		if err != nil {
			log.Println("Error creating template cache :", err)
			return
		}
	} else {
		log.Println("Using template from cache")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)

	if err != nil {
		fmt.Println("Error executing template :", err)
		return
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		"../../templates/" + t,
		"../../templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	tc[t] = tmpl

	return nil
}
