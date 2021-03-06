package authorization

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/drew138/games/database"
	"github.com/drew138/games/database/models"
)

// RefreshToken - generate a new pair authorization JWT if user ID exists
func RefreshToken(token *jwt.Token) (map[string]string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		var User models.User // user in database
		database.DBConn.First(&User, claims["ID"])
		tokenPair, err := GenerateJWT(&User)
		return tokenPair, err
	}
	return map[string]string{}, fmt.Errorf("Error: Inavlid token")
}
