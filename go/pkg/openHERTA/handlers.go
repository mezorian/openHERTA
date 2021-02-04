package openHERTA

import (
	"fmt"
	"log"
	"net/http"
)

// Home is the home page handler
func (oH *OpenHERTA) Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "../../static-files/html/event.html")
}

// About is the about page handler
func (oH *OpenHERTA) About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

// Test is the test page handler
func (oH *OpenHERTA) Test(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "test.tmpl")
}

// Test is the test page handler
func (oH *OpenHERTA) TemplateTest(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "TemplateTest.tmpl")
}

func (oH *OpenHERTA) GetHeader(w http.ResponseWriter, r *http.Request) {
	eventID := "75683"
	renderHeaderDiv(w, oH.Events[eventID])
}

func (oH *OpenHERTA) GetContentMap(w http.ResponseWriter, r *http.Request) {
	eventID := "75683"
	renderContentMapDiv(w, oH.Events[eventID])
}

func (oH *OpenHERTA) GetGuests(w http.ResponseWriter, r *http.Request) {
	eventID := "75683"
	renderGuestsDiv(w, oH.Events[eventID])
}

func (oH *OpenHERTA) GetContentBringSomething(w http.ResponseWriter, r *http.Request) {
	eventID := "75683"
	renderContentBringSomethingDiv(w, oH.Events[eventID])
}

func (oH *OpenHERTA) GetContentHeaderDetails(w http.ResponseWriter, r *http.Request) {
	eventID := "75683"
	renderContentHeaderDetailsDiv(w, oH.Events[eventID])
}

func (oH *OpenHERTA) GetContentGuest(w http.ResponseWriter, r *http.Request) {
	eventID := "75683"
	renderContentGuestDiv(w, oH.Events[eventID])
}

func (oH *OpenHERTA) GetContentDiscussion(w http.ResponseWriter, r *http.Request) {
	eventID := "75683"
	renderContentDiscussionDiv(w, oH.Events[eventID])
}

/* --- UPDATE DATA --- */
func (oH *OpenHERTA) UpdateGuestDetailsName(w http.ResponseWriter, r *http.Request) {
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	userID := r.FormValue("userID")
	sessionID := r.FormValue("sessionID")
	eventID := r.FormValue("eventID")
	fmt.Fprintf(w, "Hello World %s %s %s %s %s", firstName, lastName, userID, sessionID, eventID)
	log.Printf("Hello World %s %s %s %s %s", firstName, lastName, userID, sessionID, eventID)
}

func (oH *OpenHERTA) UpdateConfirmationStatus(w http.ResponseWriter, r *http.Request) {
	confirmationStatus := r.FormValue("ConfirmationStatus")
	userID := r.FormValue("UserID")
	sessionID := r.FormValue("SessionID")
	eventID := r.FormValue("EventID")
	fmt.Fprintf(w, "hi your status is %s, with userID %s, sessionID %s, eventID %s", confirmationStatus, userID, sessionID, eventID)
}

func (oH *OpenHERTA) UpdateBringSomeone(w http.ResponseWriter, r *http.Request) {
	numberOfPeopleThisGuestWillBring := r.FormValue("bring_someone")
	fmt.Fprintf(w, "bring %s", numberOfPeopleThisGuestWillBring)
}

func (oH *OpenHERTA) UpdateBringSomething(w http.ResponseWriter, r *http.Request) {
	category := r.FormValue("bring_something_category_select")
	name := r.FormValue("bring_something")
	fmt.Fprintf(w, "bring %s %s", category, name)
}

func (oH *OpenHERTA) UpdateDiscussion(w http.ResponseWriter, r *http.Request) {
	new_message_text_area := r.FormValue("content_discussion_new_message_text_area")
	fmt.Fprintf(w, "new message %s", new_message_text_area)
}
