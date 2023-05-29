package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Dmkk01/bed-and-breakfast/pkg/config"
	"github.com/Dmkk01/bed-and-breakfast/pkg/handlers"
	"github.com/Dmkk01/bed-and-breakfast/pkg/render"
)

const portNumber = ":11111"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s \n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
