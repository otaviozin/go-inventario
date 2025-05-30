package item

import (
	"github.com/gofiber/fiber/v2"
	itemController "github.com/otaviozin/go-inventory/controllers/item"
)

func RegisterRoutes(api fiber.Router) {
	item := api.Group("/item")

	item.Get("/", itemController.GetItem)

	item.Post("/", itemController.Create)
}
