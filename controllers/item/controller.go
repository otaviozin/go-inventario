package item

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/otaviozin/go-inventory/services/item"
	"github.com/otaviozin/go-inventory/types"
)

func Create(c *fiber.Ctx) error {
	var input types.Item
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
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

func GetById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	if idParam == "" {
		return c.Status(400).JSON(fiber.Map{"error": "ID is required"})
	}

	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid item ID"})
	}

	itemFound, err := item.GetById(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Item not found"})
	}

	return c.Status(200).JSON(itemFound)
}
