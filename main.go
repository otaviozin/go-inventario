package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/otaviozin/go-inventory/routes"
)

func main() {
	app := fiber.New()
	routes.RoutePrefix(app)
	app.Listen(":8000")
}
