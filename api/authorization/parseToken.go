package authorization

import jwt "github.com/dgrijalva/jwt-go"

// ParseJWT - retrieve data from JWT
func ParseJWT(token string) (string, error) {
	jwt.Parse(token)
}
