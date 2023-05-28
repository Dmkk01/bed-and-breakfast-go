package handlers

import (
	"net/http"

	"github.com/Dmkk01/bed-and-breakfast/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateTest(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateTest(w, "about.page.tmpl")
}
