package handler

import (
	mdl "distr-calc/model"
	"net/http"
	"text/template"
)

// settings handler - returns the settings.html template, with settings data
func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/settings.html"))
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
