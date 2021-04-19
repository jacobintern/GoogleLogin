package controllers

import (
	"net/http"
	"text/template"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./views/login.html"))
	tmpl.Execute(w, nil)
}
