package handler

import (
	"context"
	"fmt"
	"net/http"
	"text/template"
	mdl "yc/distr-calc/model"
)

// settings handler - returns the settings.html template, with settings data
func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/settings.html"))

	login := getParams(r.Context())
	fmt.Println(login)

	settings := map[string][]mdl.Setting{
		"Settings": {
			{Name: "Operation execution time +", Value: 200},
			{Name: "Operation execution time -", Value: 200},
			{Name: "Operation execution time *", Value: 200},
			{Name: "Operation execution time /", Value: 200},
			{Name: "The display time of the inactive server", Value: 200},
		},
	}
	tmpl.Execute(w, settings)
}

func getParams(ctx context.Context) string {
	fmt.Println("getParams")
	if ctx == nil {
		fmt.Println("getParams ctx nil")
		return ""
	}

	login, ok := ctx.Value("login").(string)
	if ok {
		return login
	}

	return ""
}
