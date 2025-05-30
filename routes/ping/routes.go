package ping

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(api fiber.Router) {
	ping := api.Group("/ping")

	ping.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}