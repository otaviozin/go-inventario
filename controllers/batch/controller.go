package batch

import (
	"github.com/gofiber/fiber/v2"
	"github.com/otaviozin/go-inventory/services/batch"
	"github.com/otaviozin/go-inventory/services/item"
	"github.com/otaviozin/go-inventory/types"
)

func Create(c *fiber.Ctx) error {
	var input types.Batch
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	if input.ItemID == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "ItemID is required"})
	}

	_, err := item.GetById(input.ItemID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Item not found"})
	}

	newBatch, err := batch.Create(input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error creating Batch"})
	}

	return c.Status(201).JSON(newBatch)
}
