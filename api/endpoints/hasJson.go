package endpoints

import (
	"github.com/gofiber/fiber"
)

// HasJSONBody - Send 400 response in case request doesnt have a json body
func HasJSONBody(c *fiber.Ctx) bool {
	if !c.Is("json") {
		c.Status(400).Send("Bad Request: content type json not found")
		return false
	}
	return true
}
