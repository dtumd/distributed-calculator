package handler

import (
	db "distr-calc/db"
	mdl "distr-calc/model"
	"distr-calc/parse"
	"fmt"
	"net/http"
	"strings"
	"text/template"
	"time"
)

// solve
func Solve(s string) float64 {
	p := parse.NewParser(strings.NewReader(s))
	fmt.Printf("%+v\n", p)
	stack, _ := p.Parse()
	fmt.Printf("%+v\n", stack)
	stack = parse.ShuntingYard(stack)
	fmt.Printf("%+v\n", stack)
	answer := parse.SolvePostfix(stack)
	return answer
}

// add expression handler - returns the template block with the newly added expression, as an HTMX response
func AddExpressionHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	expr := r.PostFormValue("expr-val")
	fmt.Printf("%+v\n", Solve(expr))
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	e := mdl.Expression{Uuid: "2", Status: "Calculating", Value: expr, Result: "?"}
	e = db.SaveExpressions(e)

	tmpl.ExecuteTemplate(w, "expression-list-element", e)
}
