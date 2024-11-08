package routes

import (
	"mvcpattern/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Post("/user", controller.CreateUser)
	app.Get("/user", controller.UserGetController)
	app.Patch("/user/:id", controller.UpdateUser)
}
