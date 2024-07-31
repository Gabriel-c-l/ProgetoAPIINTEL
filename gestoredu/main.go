package main

import (
	"git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.AppRoutes(app)

	app.Listen(":3000")
}
