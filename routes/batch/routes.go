package batch

import (
	"github.com/gofiber/fiber/v2"
	batchController "github.com/otaviozin/go-inventory/controllers/batch"
)

func RegisterRoutes(api fiber.Router) {
	batch := api.Group("/batch")

	batch.Post("/", batchController.Create)
}
