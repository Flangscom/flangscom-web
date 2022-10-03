package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var App fiber.App = *fiber.New()

func main() {

	App.Use(cors.New())

	App.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "Users",
		})
	})

	App.Static("/", "./flangscom-web/dist")
	App.Listen(":3000")
}
