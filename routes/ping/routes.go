package ping

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router) {
	ping := router.Group("/ping")

	ping.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}