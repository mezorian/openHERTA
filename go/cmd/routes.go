package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(handlerContext *HandlerContext) http.Handler {
	mux := chi.NewRouter()

	// define middleware here
	mux.Use(middleware.Logger)
	// TODO : find out what that's doing
	//mux.Use(NoSurf)

	// define routes
	mux.Get("/", handlerContext.Home)
	mux.Get("/about", handlerContext.About)
	mux.Get("/test", handlerContext.Test)
	mux.Get("/test2", handlerContext.TemplateTest)
	// header
	mux.Get("/getHeader", handlerContext.GetHeader)
	// right container
	mux.Get("/getContentMap", handlerContext.GetContentMap)
	mux.Get("/getContentGuest", handlerContext.GetContentGuest)
	mux.Get("/getGuests", handlerContext.GetGuests)
	mux.Get("/getContentBringSomething", handlerContext.GetContentBringSomething)
	// left container
	mux.Get("/getContentHeaderDetails", handlerContext.GetContentHeaderDetails)
	mux.Get("/getContentDiscussion", handlerContext.GetContentDiscussion)

	/* --- UPDATE DATA ---*/
	mux.Post("/updateGuestDetailsName", handlerContext.UpdateGuestDetailsName)
	mux.Post("/updateConfirmationStatus", handlerContext.UpdateConfirmationStatus)
	mux.Post("/updateBringSomeone", handlerContext.UpdateBringSomeone)

	// define file handler to serve static files
	fileServerStatic := http.FileServer(http.Dir("./static-files/"))
	mux.Handle("/static-files/*", http.StripPrefix("/static-files", fileServerStatic))

	// define file handler to serve js files
	fileServerJS := http.FileServer(http.Dir("./js/"))
	mux.Handle("/js/*", http.StripPrefix("/js", fileServerJS))

	return mux
}
