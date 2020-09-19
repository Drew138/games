package main

import (
	"log"

	"github.com/drew138/games/api"
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	loadEnv()
	app := fiber.New()

	api.ResgisterEndPoints(app)
	app.Listen(3000)

}
