package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJwt(login string) string {
	const hmacSampleSecret = "super_secret_signature"
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"nbf":   now.Unix(),                        // время, с которого токен станет валидным
		"exp":   now.Add(100 * time.Minute).Unix(), // время, с которого токен перестанет быть валидным ("протухнет")
		"iat":   now.Unix(),                        // время создания токена
	})

	tokenString, err := token.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		panic(err)
	}

	fmt.Println("token string:", tokenString)

	return tokenString
}
