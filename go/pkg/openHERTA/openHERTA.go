package openHERTA

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/mezorian/openHERTA/go/pkg/config"
)

type OpenHERTA struct {
	srv *http.Server
}

func (oH *OpenHERTA) Run(wg *sync.WaitGroup) {

	handlerContext := NewHandlerContext()

	// configure the application
	config := config.Config{
		PortNumber: ":8080",
	}

	// define http server
	oH.srv = &http.Server{
		Addr:              config.PortNumber,
		Handler:           routes(handlerContext),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	// start http server
	log.Printf("Starting HTTP server on port %s....", config.PortNumber)
	wg.Add(1) // add one goroutine to waitgroup

	go func() {
		// ListenAndServe will start a http server
		// this call will block until
		// - either : the server is panic-ing
		// - or     : the Shutdown function is called (which does a graceful shutdown of the http-server)
		error := oH.srv.ListenAndServe()

		// as soon as the blocking command is finished we need to
		// tell the waitgroup that this go-routine is finished
		// wg.Done() will do this by decremnting the wait-group counter as soon
		// as this function exits
		wg.Done()
		log.Printf("OpenHERTA HTTP Server stopped with message '%s'", error)

	}()

}

func (oH *OpenHERTA) Shutdown(wg *sync.WaitGroup) {
	log.Printf("Shutting down OpenHERTA HTTP server")

	// gracefully shutdown the http server
	error := oH.srv.Shutdown(context.Background())

	// if there are any errors, catch them
	// and print them to console
	if error != nil {
		log.Printf("Shutting down OpenHERTA HTTP Server failed with message '%s'", error)

	}

}
