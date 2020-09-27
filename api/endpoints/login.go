package endpoints

import (
	"github.com/drew138/games/api/authentication"
	"github.com/drew138/games/api/authorization"
	"github.com/drew138/games/database"
	"github.com/drew138/games/database/models"
	"github.com/gofiber/fiber/v2"
)

// LogIn - Grant access and permissions by providing jwt
func LogIn(c *fiber.Ctx) {
	if !HasJSONBody(c) {
		return
	}
	user := new(models.User) // request user
	if UnmarshalJSON(c, &user) {
		return
	}
	var User models.User // user in database
	database.DBConn.Where("email = ?", user.Email).First(&User)
	err := authentication.AssertPassword(User.Password, []byte(user.Password))
	if err != nil {
		c.Status(401).Send(err)
		return
	}
	tokenPair, err := authorization.GenerateJWT(user)
	if err != nil {
		c.Status(401).Send(err)
		return
	}
	c.Status(201)
	c.JSON(tokenPair)
}
