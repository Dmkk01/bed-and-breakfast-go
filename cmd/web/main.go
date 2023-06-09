package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"

	"github.com/Dmkk01/bed-and-breakfast/internal/config"
	"github.com/Dmkk01/bed-and-breakfast/internal/handlers"
	"github.com/Dmkk01/bed-and-breakfast/internal/models"
	"github.com/Dmkk01/bed-and-breakfast/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":11111"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	gob.Register(models.Reservation{})

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * 60 * 60
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s \n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
