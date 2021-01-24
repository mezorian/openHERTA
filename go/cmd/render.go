package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/mezorian/openHERTA/go/pkg/event"
	"github.com/mezorian/openHERTA/go/pkg/guest"
)

func renderGuestsDiv(w http.ResponseWriter, e *event.Event) {
	tmpl := "guests.tmpl"
	parsedTemplate, _ := template.ParseFiles("go/templates/" + tmpl)

	data := struct {
		NumberOfConfirmedGuests  int
		NumberOfInterestedGuests int
		NumberOfDeclinedGuests   int
		ConfirmedGuests          []*guest.Guest
		InterestedGuests         []*guest.Guest
		DeclinedGuests           []*guest.Guest
	}{
		e.GetNumberOfConfirmedGuests(),
		e.GetNumberOfInterestedGuests(),
		e.GetNumberOfDeclinedGuests(),
		e.GetConfirmedGuests(),
		e.GetInterestedGuests(),
		e.GetDeclinedGuests(),
	}

	err := parsedTemplate.Execute(w, data)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	if tmpl == "TemplateTest.tmpl" {
		fmt.Println("template test")
		parsedTemplate, _ := template.ParseFiles("go/templates/layout.tmpl", "go/templates/example.tmpl")
		err := parsedTemplate.Execute(w, nil)
		if err != nil {
			fmt.Println("error parsing template: ", err)
			return
		}
	} else {
		parsedTemplate, _ := template.ParseFiles("go/templates/" + tmpl)
		err := parsedTemplate.Execute(w, nil)
		if err != nil {
			fmt.Println("error parsing template: ", err)
			return
		}
	}

}
