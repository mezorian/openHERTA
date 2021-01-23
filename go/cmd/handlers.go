package main

import (
	"fmt"
	"net/http"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "../../static-files/html/event.html")

}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")

}

// Test is the test page handler
func Test(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "test.tmpl")

}

// Test is the test page handler
func TemplateTest(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "TemplateTest.tmpl")
}

func GetGuests(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
