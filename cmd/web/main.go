package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/0x1david/monolith-app/pkg/config"
	"github.com/0x1david/monolith-app/pkg/handlers"
	"github.com/0x1david/monolith-app/pkg/render"
)

const portNumber = ":8080"


//main is the main application function
func main() {
    var app config.AppConfig

    templateCache, err := render.CreateTemplateCache()
    if err != nil {
        log.Fatal("Cannot create template cache in main app.")
    }

    app.TemplateCache = templateCache
    app.UseCache = false

    repo := handlers.NewRepo(&app)
    handlers.NewHandlers(repo)

    render.NewTemplates(&app)

    http.HandleFunc("/", handlers.Repo.Home)
    http.HandleFunc("/about", handlers.Repo.About)
    fmt.Println(fmt.Sprintf("Starting the application on the port %s", portNumber))
    _ = http.ListenAndServe(portNumber, nil)
}
