package handler

import (
	"net/http"
	"text/template"
	o "yc/distr-calc/orchestrator"
)

// resources handler - returns the resources.html template, with film data
func ResourcesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/resources.html"))
	tmpl.Execute(w, o.GetResources())
}
