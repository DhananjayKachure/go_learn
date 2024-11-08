package main

import (
	"formbuilder/connection"
	"formbuilder/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	connection.ConnectDb()
	routes.UserRoutes(app)

	app.Listen(":3000")
}
