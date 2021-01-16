package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	// define middleware here
	// ...

	// define routes
	mux.Get("/", Home)
	mux.Get("/about", About)

	return mux
}
