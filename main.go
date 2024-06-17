package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/naol86/Go/fiber/bookstore/pkg/routes"
)

func main() {
	app := fiber.New()

	// enable cors for all origins for now
	app.Use(cors.New())

	routes.Routes(app)
	app.Listen(":8000")

}