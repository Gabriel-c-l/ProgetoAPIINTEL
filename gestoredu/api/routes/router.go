package routes

import (
	"git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/api/controller"
	"github.com/gofiber/fiber/v2"
)

func AppRoutes(app *fiber.App) {
	usersController := controller.NewUsersController()

	v1 := app.Group("/v1")
	v1.Post("/user", usersController.Create)
	v1.Patch("/user/:id", usersController.Patch)
	v1.Delete("/user/:id", usersController.Delete)
}
