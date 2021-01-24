package main

import (
	"net/http"

	"github.com/mezorian/openHERTA/go/pkg/event"
)

type HandlerContext struct {
	Events []*event.Event
}

// constructor of HandlerContext
func NewHandlerContext() *HandlerContext {
	// create empty HandlerContext object
	result := new(HandlerContext)

	// create new dummy event object
	dummyEvent := new(event.Event)
	dummyEvent.CreateDummyData()

	// create empty slice of events and add dummy event to it
	result.Events = make([]*event.Event, 0)
	result.Events = append(result.Events, dummyEvent)

	return result

}

// Home is the home page handler
func (c *HandlerContext) Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "../../static-files/html/event.html")
}

// About is the about page handler
func (c *HandlerContext) About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

// Test is the test page handler
func (c *HandlerContext) Test(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "test.tmpl")
}

// Test is the test page handler
func (c *HandlerContext) TemplateTest(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "TemplateTest.tmpl")
}

func (c *HandlerContext) GetGuests(w http.ResponseWriter, r *http.Request) {
	for _, g := range c.Events[0].Guests {
		println(g.FirstName)
		println(g.LastName)

	}
	renderGuestsDiv(w, c.Events[0].Guests)
}
