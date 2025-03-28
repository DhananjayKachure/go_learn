package routes

import (
	"formbuilder/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	app.Post("/user", controller.RegisterUser)
	app.Post("/user/login", controller.LoginUser)
	app.Post("/user/forgot", controller.ForgetPassword)
	app.Get("/news", controller.GetUser)
	app.Get("/sectionlisting",controller.GetsectionListing)
	// app.Patch("/user/:id", controller.UpdateUser)
	app.Put("/user/changepassword", controller.ChangePassword)
	app.Put("/user/:id", controller.ChangeForgetPassword)
}
