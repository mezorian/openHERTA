package main

import (
	"fmt"
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

func (c *HandlerContext) GetHeader(w http.ResponseWriter, r *http.Request) {
	renderHeaderDiv(w, c.Events[0])
}

func (c *HandlerContext) GetContentMap(w http.ResponseWriter, r *http.Request) {
	renderContentMapDiv(w, c.Events[0])
}

func (c *HandlerContext) GetGuests(w http.ResponseWriter, r *http.Request) {
	renderGuestsDiv(w, c.Events[0])
}

func (c *HandlerContext) GetContentBringSomething(w http.ResponseWriter, r *http.Request) {
	renderContentBringSomethingDiv(w, c.Events[0])
}

func (c *HandlerContext) GetContentHeaderDetails(w http.ResponseWriter, r *http.Request) {
	renderContentHeaderDetailsDiv(w, c.Events[0])
}

func (c *HandlerContext) GetContentGuest(w http.ResponseWriter, r *http.Request) {
	renderContentGuestDiv(w, c.Events[0])
}

func (c *HandlerContext) GetContentDiscussion(w http.ResponseWriter, r *http.Request) {
	renderContentDiscussionDiv(w, c.Events[0])
}

/* --- UPDATE DATA --- */
func (c *HandlerContext) UpdateGuestDetailsName(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	fmt.Fprintf(w, "Hello World %s %s", firstName, lastName)
}

func (c *HandlerContext) UpdateConfirmationStatus(w http.ResponseWriter, r *http.Request) {
	confirmationStatus := r.FormValue("ConfirmationStatus")
	fmt.Fprintf(w, "hi %s", confirmationStatus)
}

func (c *HandlerContext) UpdateBringSomeone(w http.ResponseWriter, r *http.Request) {
	numberOfPeopleThisGuestWillBring := r.FormValue("bring_someone")
	fmt.Fprintf(w, "bring %s", numberOfPeopleThisGuestWillBring)
}

func (c *HandlerContext) UpdateBringSomething(w http.ResponseWriter, r *http.Request) {
	category := r.FormValue("bring_something_category_select")
	name := r.FormValue("bring_something")
	fmt.Fprintf(w, "bring %s %s", category, name)
}

func (c *HandlerContext) UpdateDiscussion(w http.ResponseWriter, r *http.Request) {
	new_message_text_area := r.FormValue("content_discussion_new_message_text_area")
	fmt.Fprintf(w, "new message %s", new_message_text_area)
}
