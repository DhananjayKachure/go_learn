package main

import (
	"formbuilder/connection"
	"formbuilder/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allows all origins, you can specify domains like "http://example.com"
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type,Authorization",
	}))
	connection.ConnectDb()
	routes.UserRoutes(app)

	app.Listen(":4000")
}
