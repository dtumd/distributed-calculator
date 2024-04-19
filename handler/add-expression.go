package handler

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
	orch "yc/distr-calc/orchestrator"
)

// add expression handler - returns the template block with the newly added expression, as an HTMX response
func AddExpressionHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	expr := r.PostFormValue("expr-val")

	login := getParams(r.Context())

	fmt.Println("AddExpressionHandler, user login: " + login)

	e := orch.AddExpression(expr, login)

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.ExecuteTemplate(w, "expression-list-element", e)
}
