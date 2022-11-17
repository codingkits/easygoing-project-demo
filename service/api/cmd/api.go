package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Use(gloM)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})
	api := app.Group("/api")
	api.Get("/device/search", func(c *fiber.Ctx) error {
		return c.SendString("api arrived")
	})
	app.Listen(":3000")
}

func gloM(c *fiber.Ctx) error {
	c.Set("GSet", "1")
	return c.Next()
}
