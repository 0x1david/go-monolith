package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/0x1david/monolith-app/pkg/config"
	"github.com/0x1david/monolith-app/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
    app = a
}

func AddDefaultData(templateData *models.TemplateData) *models.TemplateData{

    return templateData
}

func RenderTemplate(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {
    var templateCache map[string]*template.Template
    var err error

    if app.UseCache{
        templateCache = app.TemplateCache
    } else {
        templateCache, err = CreateTemplateCache()
        if err != nil {
            log.Fatal("Could not create new template cache.")
        }
    }

    template, ok := templateCache[tmpl]
    if !ok {
        log.Fatal("Could not get template from template cache.")
    }

    buf := new(bytes.Buffer)
    
    templateData = AddDefaultData(templateData)

    err = template.Execute(buf, templateData)
    if err != nil {
        log.Println(err)
    }

    _, err = buf.WriteTo(w)
    if err != nil {
        log.Println(err)
    }
}

func CreateTemplateCache() (map[string]*template.Template, error) {
    myCache := map[string]*template.Template{}

    pages, err := filepath.Glob("./templates/*.page.tmpl")
    if err != nil {
        return myCache, err
    }

    for _, page := range pages {
        name := filepath.Base(page)
        ts, err := template.New(name).ParseFiles(page)
        if err != nil {
            return myCache, err
        }

        matches, err := filepath.Glob("./templates/*.layout.tmpl")
        if err != nil {
            return myCache, err
        }
        if len(matches) > 0 {
            ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
            if err != nil {
                return myCache, err
            }
        }
        myCache[name] = ts
    } 
    return myCache, nil
}

