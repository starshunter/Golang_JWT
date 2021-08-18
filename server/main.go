package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/starshunter/Golang_JWT/database"
	"github.com/starshunter/Golang_JWT/routes"
)

func handler(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8080")
}
