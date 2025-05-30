package item

import (
	"github.com/gofiber/fiber/v2"
	"github.com/otaviozin/go-inventory/services/item"
	"github.com/otaviozin/go-inventory/types"
)

func Create(c *fiber.Ctx) error {
	var input types.Item
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "JSON inválido"})
	}

	if input.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid name"})
	}

	newItem, err := item.Create(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error creating item"})
	}

	return c.Status(201).JSON(newItem)
}
