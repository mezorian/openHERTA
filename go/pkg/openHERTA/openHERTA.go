package openHERTA

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/mezorian/openHERTA/go/pkg/config"
	"github.com/mezorian/openHERTA/go/pkg/event"
	"github.com/mezorian/openHERTA/go/pkg/guest"
)

type OpenHERTA struct {
	srv    *http.Server
	Events map[string]*event.Event
}

func (oH *OpenHERTA) Run(wg *sync.WaitGroup) {

	// configure the application
	config := config.Config{
		PortNumber: ":8080",
	}

	// define http server
	oH.srv = &http.Server{
		Addr:              config.PortNumber,
		Handler:           routes(oH),
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

func (oH *OpenHERTA) Shutdown() {
	log.Printf("Shutting down OpenHERTA HTTP server")

	// gracefully shutdown the http server
	error := oH.srv.Shutdown(context.Background())

	// if there are any errors, catch them
	// and print them to console
	if error != nil {
		log.Printf("Shutting down OpenHERTA HTTP Server failed with message '%s'", error)
	}
}

/* --- INITIALIZE WEBSITE WITH TESTING DATA ---  */

func (oH *OpenHERTA) CreateDummyData() {

	// initialize Events map
	oH.Events = make(map[string]*event.Event)

	eventID := "75683"
	dummyEvent := new(event.Event)
	oH.Events[eventID] = dummyEvent

	oH.Events[eventID].Name = "Lorem ipsum Grillabend"
	oH.Events[eventID].InfoTitle = "Lorem ipsum Grillabend"
	oH.Events[eventID].Date = "Samstag, 20. Juni 2021"
	oH.Events[eventID].Time = "18:00 Uhr"
	oH.Events[eventID].AddressString = "Hinten im Garten"
	oH.Events[eventID].AddressGPSLat = 52.297652
	oH.Events[eventID].AddressGPSLong = -10.0383357
	// Todo
	oH.Events[eventID].Description = ""

	dummyGuest := new(guest.Guest)
	dummyGuest.FirstName = "John"
	dummyGuest.LastName = "Stone"
	dummyGuest.Confirm()
	oH.Events[eventID].AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Ponnappa"
	dummyGuest.LastName = "Priya"
	dummyGuest.Confirm()
	oH.Events[eventID].AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Mia"
	dummyGuest.LastName = "Wong"
	dummyGuest.Confirm()
	dummyGuest.BringSomeone(2)
	oH.Events[eventID].AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Peter"
	dummyGuest.LastName = "Stanbridge"
	dummyGuest.Confirm()
	oH.Events[eventID].AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Natalie"
	dummyGuest.LastName = "Lee-Walsh"
	dummyGuest.Confirm()
	oH.Events[eventID].AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Ang"
	dummyGuest.LastName = "Li"
	dummyGuest.BringSomeone(2)
	dummyGuest.Confirm()
	oH.Events[eventID].AddGuest(dummyGuest)

	// interested guests

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Nguta"
	dummyGuest.LastName = "Ithya"
	dummyGuest.Interest()
	oH.Events[eventID].AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Tamzyn"
	dummyGuest.LastName = "French"
	dummyGuest.Interest()
	oH.Events[eventID].AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Salome"
	dummyGuest.LastName = "Simoes"
	dummyGuest.Interest()
	dummyGuest.BringSomeone(1)
	oH.Events[eventID].AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Trevor"
	dummyGuest.LastName = "Virtue"
	dummyGuest.Interest()
	dummyGuest.BringSomeone(3)
	oH.Events[eventID].AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Tarryn"
	dummyGuest.LastName = "Campbell-Gillies"
	dummyGuest.Interest()
	dummyGuest.BringSomeone(2)
	oH.Events[eventID].AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Eugenia"
	dummyGuest.LastName = "Anders"
	dummyGuest.Interest()
	oH.Events[eventID].AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Andrew"
	dummyGuest.LastName = "Kazantzis"
	dummyGuest.Interest()
	dummyGuest.BringSomeone(2)
	oH.Events[eventID].AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Verona"
	dummyGuest.LastName = "Blair"
	dummyGuest.Interest()
	oH.Events[eventID].AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Jane"
	dummyGuest.LastName = "Meldrum"
	dummyGuest.Interest()
	dummyGuest.BringSomeone(2)
	oH.Events[eventID].AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Maureen"
	dummyGuest.LastName = "M. Smith"
	dummyGuest.Interest()
	oH.Events[eventID].AddGuest(dummyGuest)

	// declined guests

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Desiree"
	dummyGuest.LastName = "Burch"
	dummyGuest.Decline()
	oH.Events[eventID].AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Daly"
	dummyGuest.LastName = "Harry"
	dummyGuest.Decline()
	oH.Events[eventID].AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Hayman"
	dummyGuest.LastName = "Andrews"
	dummyGuest.Decline()
	oH.Events[eventID].AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Ruveni"
	dummyGuest.LastName = "Ellawala"
	dummyGuest.Decline()
	oH.Events[eventID].AddGuest(dummyGuest)
}
