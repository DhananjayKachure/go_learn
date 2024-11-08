package main

import (
	"log"
	"mvcpattern/connection"
	"mvcpattern/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	connection.ConnectDb()
	routes.UserRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
