package handlers

import (
	"net/http"

	"github.com/Dmkk01/bed-and-breakfast/pkg/config"
	"github.com/Dmkk01/bed-and-breakfast/pkg/models"
	"github.com/Dmkk01/bed-and-breakfast/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// Create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	stringMap["test"] = "Hello, again."

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
