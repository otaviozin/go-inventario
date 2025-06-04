package box

import (
	"github.com/gofiber/fiber/v2"
	boxController "github.com/otaviozin/go-inventory/controllers/box"
)

func RegisterRoutes(api fiber.Router) {
	box := api.Group("/box")

	box.Post("/", boxController.Create)
}
