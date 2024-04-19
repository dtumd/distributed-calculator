package main

import (
	"fmt"
	"log"
	"net/http"
	db "yc/distr-calc/db"
	h "yc/distr-calc/handler"
)

// go run cmd/orchestrator/main.go
func main() {
	fmt.Println("Distributed Calculator app...")

	db.Init()

	mux := http.NewServeMux()

	// define handlers
	mux.Handle("/", h.CheckJwtMiddleware(http.HandlerFunc(h.IndexHandler)))
	mux.Handle("/add-expression", h.CheckJwtMiddleware(http.HandlerFunc(h.AddExpressionHandler)))
	mux.Handle("/settings", h.CheckJwtMiddleware(http.HandlerFunc(h.SettingsHandler)))
	mux.Handle("/resources", h.CheckJwtMiddleware(http.HandlerFunc(h.ResourcesHandler)))

	mux.HandleFunc("/login", h.LoginHandler)
	mux.HandleFunc("/register", h.RegisterHandler)

	mux.HandleFunc("/api/v1/login", h.LoginFormHandler)
	mux.HandleFunc("/api/v1/register", h.RegisterFormHandler)

	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}
