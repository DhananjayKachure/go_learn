package main

import (
	"log"
	"noteapp/config"
	"noteapp/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.ConnectDb()
	app := fiber.New()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
