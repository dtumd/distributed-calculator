package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtc, err := r.Cookie("jwt")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
		}

		fmt.Println(jwtc)

		v := 123

		rcopy := r.WithContext(context.WithValue(r.Context(), "uid", v))

		next.ServeHTTP(w, rcopy)
	})
}

func CheckJwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtc, err := r.Cookie("jwt")
		if err != nil || jwtc == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			next.ServeHTTP(w, r)
			return
		}

		fmt.Println(jwtc)

		const hmacSampleSecret = "super_secret_signature"

		tokenFromString, err := jwt.Parse(jwtc.Value, func(token *jwt.Token) (interface{}, error) {
			fmt.Println("jwt.Parse")
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				//panic(fmt.Errorf("unexpected signing method: %v", token.Header["alg"]))
				fmt.Println("not ok")
				// http.Redirect(w, r, "/login", http.StatusSeeOther)
				// next.ServeHTTP(w, r)
				return nil, fmt.Errorf("")
			}

			return []byte(hmacSampleSecret), nil
		})

		if err != nil {
			//log.Fatal(err)
			fmt.Println("err not nil")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			next.ServeHTTP(w, r)
			return
		}

		claims, ok := tokenFromString.Claims.(jwt.MapClaims)
		fmt.Println("user login: ", claims["login"])

		if !ok {
			//panic(err)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			next.ServeHTTP(w, r)
			return
		}

		rcopy := r.WithContext(context.WithValue(r.Context(), "login", claims["login"]))

		next.ServeHTTP(w, rcopy)
	})
}
