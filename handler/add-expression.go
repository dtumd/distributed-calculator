package handler

import (
	db "distr-calc/db"
	mdl "distr-calc/model"
	"net/http"
	"text/template"
	"time"
)

// add expression handler - returns the template block with the newly added expression, as an HTMX response
func AddExpressionHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	expr := r.PostFormValue("expr-val")
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	e := mdl.Expression{Uuid: "2", Status: "Calculating", Value: expr, Result: "?"}
	db.SaveExpressions(e)

	tmpl.ExecuteTemplate(w, "expression-list-element", e)
}
