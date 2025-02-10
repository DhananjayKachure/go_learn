package routes

import (
	"noteapp/controllers"
	"noteapp/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/app")
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)
	api.Use(middleware.AuthMiddleware)
	api.Post("/notes", controllers.CreateNote)
	api.Get("/notes", controllers.GetNotes)

}
