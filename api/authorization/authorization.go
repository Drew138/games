package api

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateJWT return jwt token
func GenerateJWT() (string, error) {
	var secretKey = os.Getenv("SECRET_KEY")
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	// claims["user"] =
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
