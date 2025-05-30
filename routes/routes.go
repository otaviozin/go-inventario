package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/otaviozin/go-inventario/routes/ping"
)

func RoutePrefix(app *fiber.App) {
	api := app.Group("/api")

	ping.RegisterRoutes(api)
}