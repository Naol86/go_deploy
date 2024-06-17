package main

import "github.com/gofiber/fiber/v2";

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}


func main() {

	app := fiber.New()

	app.Get("/", Hello)

	app.Listen("0.0.0.0:")

}