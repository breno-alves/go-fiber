package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		msg := "Hello world"
		return c.SendString(msg)
	})

	log.Fatal(app.Listen(":3000"))
}
