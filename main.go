package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/naol86/Go/fiber/bookstore/pkg/routes"
)

func main() {
	app := fiber.New()

	// enable cors for all origins for now
	app.Use(cors.New())
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	routes.Routes(app)
	log.Fatal(app.Listen(":" + port))

}