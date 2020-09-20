package endpoints

import (
	"encoding/json"

	"github.com/gofiber/fiber"
)

// UnmarshalJSON - apply json body of a request context to an especified model
func UnmarshalJSON(c *fiber.Ctx, model interface{}) bool {
	err := json.Unmarshal([]byte(c.Body()), model)
	if err != nil {
		c.Status(400).Send("Bad Request: invalid content field")
		return true
	}
	return false
}
