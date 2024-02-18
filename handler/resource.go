package handler

import (
	o "distr-calc/orchestrator"
	"net/http"
	"text/template"
)

// resources handler - returns the resources.html template, with film data
func ResourcesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/resources.html"))
	tmpl.Execute(w, o.GetResources())
}
