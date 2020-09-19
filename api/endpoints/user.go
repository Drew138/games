package endpoints

import (
	"encoding/json"

	"github.com/drew138/games/api/authentication"
	"github.com/drew138/games/database"
	"github.com/drew138/games/database/models"
	"github.com/gofiber/fiber"
)

// CreateUser add new user to database
func CreateUser(c *fiber.Ctx) {

	if !c.Is("json") {
		c.Status(400).Send("Bad Request: content type json not found")
		return
	}
	user := new(models.User)
	err := json.Unmarshal([]byte(c.Body()), &user)
	if err != nil {
		c.Status(400).Send("Bad Request: invalid content field")
		return
	}
	validationError := authentication.ValidatePassword(user.Password)
	if validationError != nil {
		c.Status(400).Send(validationError)
		return
	}
	user.Password = authentication.HashGenerator([]byte(user.Password))
	dbError := database.DBConn.Create(&user)
	if dbError != nil {

	}
	userMap := map[string]string{
		"email":   user.Email,
		"name":    user.Name,
		"surname": user.Surname,
	}

	c.Status(200)
	c.JSONP(userMap)
}
