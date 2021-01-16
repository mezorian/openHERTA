package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	// define middleware here
	mux.Use(middleware.Logger)

	// define routes
	mux.Get("/", Home)
	mux.Get("/about", About)

	return mux
}
