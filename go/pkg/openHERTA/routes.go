package openHERTA

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(oH *OpenHERTA) http.Handler {
	mux := chi.NewRouter()

	// define middleware here
	mux.Use(middleware.Logger)
	// TODO : find out what that's doing
	//mux.Use(NoSurf)

	// define routes
	mux.Get("/", oH.Home)
	mux.Get("/about", oH.About)
	mux.Get("/test", oH.Test)
	mux.Get("/test2", oH.TemplateTest)
	// header
	mux.Get("/getHeader", oH.GetHeader)
	// right container
	mux.Get("/getContentMap", oH.GetContentMap)
	mux.Get("/getContentGuest", oH.GetContentGuest)
	mux.Get("/getGuests", oH.GetGuests)
	mux.Get("/getContentBringSomething", oH.GetContentBringSomething)
	// left container
	mux.Get("/getContentHeaderDetails", oH.GetContentHeaderDetails)
	mux.Get("/getContentDiscussion", oH.GetContentDiscussion)

	/* --- UPDATE DATA ---*/
	mux.Post("/updateGuestDetailsName", oH.UpdateGuestDetailsName)
	mux.Post("/updateConfirmationStatus", oH.UpdateConfirmationStatus)
	mux.Post("/updateBringSomeone", oH.UpdateBringSomeone)
	mux.Post("/updateBringSomething", oH.UpdateBringSomething)
	mux.Post("/updateDiscussion", oH.UpdateDiscussion)

	// define file handler to serve static files
	fileServerStatic := http.FileServer(http.Dir("./static-files/"))
	mux.Handle("/static-files/*", http.StripPrefix("/static-files", fileServerStatic))

	// define file handler to serve js files
	fileServerJS := http.FileServer(http.Dir("./js/"))
	mux.Handle("/js/*", http.StripPrefix("/js", fileServerJS))

	return mux
}
