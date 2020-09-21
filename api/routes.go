package api

import (
	"github.com/drew138/games/api/endpoints"
	"github.com/gofiber/fiber"
)

// ResgisterEndPoints applies specified routes to fiber app
func ResgisterEndPoints(app *fiber.App) {
	// GET Endpoints
	// app.Get("/")
	// POST Endpoints
	app.Post("/api/v1/register", endpoints.CreateUser)
	app.Post("/api/v1/login", endpoints.LogIn)
	app.Post("/api/v1/logout", endpoints.Logout)
	//PUT Endpoints

	//PATCH Endpoints

}
