package handler

import (
	"fmt"
	"net/http"
	"text/template"
	auth "yc/distr-calc/auth"
	db "yc/distr-calc/db"
)

// register handler - returns the register.html template, with expressions data
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/register.html"))
	tmpl.Execute(w, nil)
}

func RegisterFormHandler(w http.ResponseWriter, r *http.Request) {
	rn := r.PostFormValue("register-name")
	rl := r.PostFormValue("register-login")
	rp := r.PostFormValue("register-password")
	rrp := r.PostFormValue("register-repeat-password")

	fmt.Println(rn, rl, rp, rrp)

	if rp != rrp {
		fmt.Println("error passwords not equal")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	}

	err := db.CreateUser(rn, rl, rp)

	if err != nil {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	}

	// create jwt
	jwt := auth.CreateJwt(rl)

	cookie := &http.Cookie{
		Name:  "jwt",
		Value: jwt,
		Path:  "/",
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
