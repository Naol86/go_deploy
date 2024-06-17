package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
);

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}


func main() {

	app := fiber.New()

	app.Get("/", Hello)
	app.Get("/env", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! from env" + os.Getenv("ENV"))
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen("0.0.0.0:" + port))

}