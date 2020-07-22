package helper

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func Token(authorizationHeader string) jwt.MapClaims {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	var key = os.Getenv("KEY")
	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	return claims
}

func DataToken(authorizationHeader string) jwt.MapClaims {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	var key = os.Getenv("KEY")
	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	claims := jwt.MapClaims{}
	token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if token.Valid {
		return claims
	}
	return claims
}
