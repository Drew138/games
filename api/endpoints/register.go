package endpoints

import (
	"fmt"

	"github.com/drew138/games/api/authentication"
	"github.com/drew138/games/database"
	"github.com/drew138/games/database/models"
	"github.com/gofiber/fiber/v2"
)

// CreateUser add new user to database
func CreateUser(c *fiber.Ctx) error {

	if !HasJSONBody(c) {
		return fmt.Errorf("Body does not contain JSON format")
	}
	user := new(models.User)
	if UnmarshalJSON(c, &user) {
		return fmt.Errorf("Invalid user properties")
	}
	validationError := authentication.ValidatePassword(user.Password)
	if validationError != nil {
		c.Status(400).Send([]byte(validationError.Error()))
		return validationError
	}
	user.Password = authentication.HashGenerator([]byte(user.Password))
	dbError := database.DBConn.Create(&user).Error
	if dbError != nil {
		c.Status(500).Send([]byte(dbError.Error()))
		return dbError
	}
	userMap := map[string]interface{}{
		"email":   user.Email,
		"name":    user.Name,
		"surname": user.Surname,
		"isAdmin": user.IsAdmin, //TODO remove this field
	}
	c.Status(201)
	c.JSON(userMap) // convert to json and send response
	return nil
}
