package handler

import (
	db "distr-calc/db"
	"net/http"
	"text/template"
)

// index handler - returns the index.html template, with expressions data
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, db.GetExpressions())
}
