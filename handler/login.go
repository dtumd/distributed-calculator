package handler

import (
	"fmt"
	"net/http"
	"text/template"
	auth "yc/distr-calc/auth"
	db "yc/distr-calc/db"
)

// login handler - returns the login.html template, with expressions data
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/login.html"))
	tmpl.Execute(w, nil) //, db.GetExpressions())
}

func LoginFormHandler(w http.ResponseWriter, r *http.Request) {
	login := r.PostFormValue("login")
	password := r.PostFormValue("password")

	fmt.Println(login, password)

	err := db.CheckPassword(login, password)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	// create jwt
	jwt := auth.CreateJwt(login)

	cookie := &http.Cookie{
		Name:  "jwt",
		Value: jwt,
		Path:  "/",
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
