package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type MyClaims struct {
	jwt.StandardClaims
	UserId string
	Email  string
}

func MakeToken(id, email string) string {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	var key = os.Getenv("KEY")
	var APPLICATION_NAME = os.Getenv("APPLICATION_NAME")

	// 1bulan
	var LOGIN_EXPIRATION_DURATION = time.Duration(1*24*30) * time.Hour
	// var LOGIN_EXPIRATION_DURATION = time.Duration(os.Getenv("EXPIRATION")) * time.Hour
	var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
	var JWT_SIGNATURE_KEY = []byte(key)
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Email:  email,
		UserId: id,
	}
	token := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)
	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return ""
	}
	return signedToken
}
