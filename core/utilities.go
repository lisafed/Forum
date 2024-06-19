package core

import (
	"html/template"
	"log"
	"net/http"
)

func LogPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("Design/style/login.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(w, nil) // Note: typically we pass some data to the template, here it's nil.
	if err != nil {
		log.Fatal(err)
	}
}
