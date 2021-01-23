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
	// TODO : find out what that's doing
	//mux.Use(NoSurf)

	// define routes
	mux.Get("/", Home)
	mux.Get("/about", About)
	mux.Get("/test", Test)
	mux.Get("/test2", TemplateTest)
	mux.Get("/getGuests", GetGuests)

	// define file handler to serve static files
	fileServerStatic := http.FileServer(http.Dir("./static-files/"))
	mux.Handle("/static-files/*", http.StripPrefix("/static-files", fileServerStatic))

	// define file handler to serve js files
	fileServerJS := http.FileServer(http.Dir("./js/"))
	mux.Handle("/js/*", http.StripPrefix("/js", fileServerJS))

	return mux
}
