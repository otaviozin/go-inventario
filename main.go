package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/otaviozin/go-inventory/database"
	"github.com/otaviozin/go-inventory/routes"
)

func main() {
	app := fiber.New()

	database.Connect()
	routes.RoutePrefix(app)

	app.Listen(":8000")
}
