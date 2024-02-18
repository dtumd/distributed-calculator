package main

import (
	db "distr-calc/db"
	h "distr-calc/handler"
	"distr-calc/parse"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// solve
func Solve(s string) float64 {
	p := parse.NewParser(strings.NewReader(s))
	stack, _ := p.Parse()
	stack = parse.ShuntingYard(stack)
	answer := parse.SolvePostfix(stack)
	return answer
}

func main() {
	fmt.Println("Distributed Calculator app...")

	Init()

	// define handlers
	http.HandleFunc("/", h.IndexHandler)
	http.HandleFunc("/add-expression/", h.AddExpressionHandler)
	http.HandleFunc("/settings", h.SettingsHandler)
	http.HandleFunc("/resources", h.ResourcesHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func Init() {
	db.InitExpressions()
	db.InitSettings()
}
