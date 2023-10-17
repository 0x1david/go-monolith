package main

import (
	"net/http"

	"github.com/0x1david/monolith-app/pkg/config"
	"github.com/0x1david/monolith-app/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)
func routes(app *config.AppConfig) http.Handler {
    mux := chi.NewRouter() 
    
    mux.Use(middleware.Recoverer)
    mux.Use(NoSurf)

    mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
    mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
    return mux
}



