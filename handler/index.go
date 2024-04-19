package handler

import (
	"fmt"
	"net/http"
	"text/template"
	db "yc/distr-calc/db"
)

// index handler - returns the index.html template, with expressions data
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	login := getParams(r.Context())
	fmt.Println(login)

	tmpl.Execute(w, db.GetExpressions())
}
