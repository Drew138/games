package endpoints

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

// SolveEquation - parse image, request solution from wolfram alpha api
func SolveEquation(c *fiber.Ctx) error {
	// TODO modify function to send request to wolfram api
	apiKey := os.Getenv("WOLFRAM_API_KEY")
	url := "http://api.wolframalpha.com/v2/query?appid=" + apiKey // TODO verify endpoint
	// https://products.wolframalpha.com/show-steps-api/documentation/
	err := proxy.Do(c, url)
	if err != nil {
		return err
	}
	// Remove Server header from response
	c.Response().Header.Del(fiber.HeaderServer)
	return nil
}
