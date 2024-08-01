package main

import (
	"git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Configuração das rotas
	routes.AppRoutes(app)

	// Inicia o servidor na porta 3000
	app.Listen(":3000")
}
