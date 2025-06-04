package box

import (
	"github.com/gofiber/fiber/v2"
	"github.com/otaviozin/go-inventory/services/box"
	"github.com/otaviozin/go-inventory/types"
)

func Create(c *fiber.Ctx) error {
	var boxes []types.Box

	if err := c.BodyParser(&boxes); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if len(boxes) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No boxes provided",
		})
	}

	if err := box.Create(boxes); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create boxes",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Boxes created successfully",
	})
}
