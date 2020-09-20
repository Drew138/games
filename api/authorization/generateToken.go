package authorization

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/drew138/games/database/models"
)

// GenerateJWT return jwt token
func GenerateJWT(user *models.User) (string, error) {
	var secretKey = os.Getenv("SECRET_KEY")
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["surname"] = user.Surname
	claims["isAdmin"] = user.IsAdmin
	claims["exp"] = time.Now().Add(time.Hour * 720).Unix() // 30 days

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		// fmt.Errorf("Error: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
