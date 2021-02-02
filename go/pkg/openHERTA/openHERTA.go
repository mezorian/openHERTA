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
		err := oH.srv.ListenAndServe()
		log.Fatal(err)
		wg.Done() // mark go-routine in waitgroup as finished
	}()

}

func (oH *OpenHERTA) Shutdown() {
	error := oH.srv.Shutdown(context.Background())

	// catch all errors
	if error != nil {
		log.Printf("listen:%+s\n", error)
	}
}
