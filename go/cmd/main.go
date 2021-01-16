package main

import (
	"log"
	"net/http"
	"time"

	"github.com/mezorian/openHERTA/go/pkg/config"
)

func main() {
	// configure the application
	config := config.Config{
		PortNumber: ":8080",
	}

	// define http server
	srv := &http.Server{
		Addr:              config.PortNumber,
		Handler:           routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	// start http server
	log.Printf("Starting HTTP server on port %s....", config.PortNumber)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
