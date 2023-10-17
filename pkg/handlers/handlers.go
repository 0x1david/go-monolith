package handlers

import (
	"net/http"

	"github.com/0x1david/monolith-app/pkg/config"
	"github.com/0x1david/monolith-app/pkg/models"
	"github.com/0x1david/monolith-app/pkg/render"
)


var Repo *Repository

// Repository is the repository type
type Repository struct {
    App *config.AppConfig
}

// NewRepo creates a new repository 
func NewRepo(a *config.AppConfig) *Repository {
    return &Repository{
        App: a,
    }
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
    Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "./templates/home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

    stringMap := make(map[string]string)
    stringMap["test"] = "Hello, again"
	render.RenderTemplate(w, "./templates/about.page.tmpl", &models.TemplateData{})
}
