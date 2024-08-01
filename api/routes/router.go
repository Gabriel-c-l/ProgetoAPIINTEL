package routes

import (
	"git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/api/controller"
	"git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/database"
	"github.com/gofiber/fiber/v2"
)

func AppRoutes(app *fiber.App) {
	db := database.DbConfig()
	usersController := controller.NewUsersController(db)

	v1 := app.Group("/v1")
	v1.Post("/user", usersController.CreateUser)
	v1.Get("/users", usersController.GetAllUsers)
	v1.Get("/user/:id", usersController.GetUserByID)
	v1.Delete("/user/:id", usersController.DeleteUserByID)
	v1.Patch("/user/:id", usersController.UpdateUserByID)
}
