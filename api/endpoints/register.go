package endpoints

import (
	"github.com/drew138/games/api/authentication"
	"github.com/drew138/games/database"
	"github.com/drew138/games/database/models"
	"github.com/gofiber/fiber"
)

// CreateUser add new user to database
func CreateUser(c *fiber.Ctx) {

	if !HasJSONBody(c) {
		return
	}
	user := new(models.User)
	if UnmarshalJSON(c, &user) {
		return
	}
	validationError := authentication.ValidatePassword(user.Password)
	if validationError != nil {
		c.Status(400).Send(validationError)
		return
	}
	user.Password = authentication.HashGenerator([]byte(user.Password))
	dbError := database.DBConn.Create(&user).Error
	if dbError != nil {
		c.Status(500).Send(dbError)
		return
	}
	userMap := map[string]interface{}{
		"email":   user.Email,
		"name":    user.Name,
		"surname": user.Surname,
		"isAdmin": user.IsAdmin, //TODO remove this field
	}

	c.Status(200)
	c.JSON(userMap) // convert to json and send response
}
