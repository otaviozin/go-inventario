package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/otaviozin/go-inventory/routes/batch"
	"github.com/otaviozin/go-inventory/routes/item"
	"github.com/otaviozin/go-inventory/routes/ping"
)

func RoutePrefix(app *fiber.App) {
	api := app.Group("/api")

	ping.RegisterRoutes(api)
	item.RegisterRoutes(api)
	batch.RegisterRoutes(api)
}
