package main

import (
	db "distr-calc/db"
	h "distr-calc/handler"
	"fmt"
	"log"
	"net/http"
)

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
