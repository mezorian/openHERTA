package main

import (
	"fmt"
	"html/template"
	"net/http"
)

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
