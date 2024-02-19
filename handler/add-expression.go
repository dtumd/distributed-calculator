package handler

import (
	orch "distr-calc/orchestrator"
	"net/http"
	"text/template"
	"time"
)

// add expression handler - returns the template block with the newly added expression, as an HTMX response
func AddExpressionHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	expr := r.PostFormValue("expr-val")
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	e := orch.AddExpression(expr)

	tmpl.ExecuteTemplate(w, "expression-list-element", e)
}
