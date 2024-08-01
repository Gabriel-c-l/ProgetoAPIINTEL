package controller

import (
	"net/http"
	"git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/api/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UsersController struct{}

func NewUsersController() *UsersController {
	return &UsersController{}
}

// Create
func (uc *UsersController) CreateUser(c *fiber.Ctx) error {
	var user entities.Users
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse body",
		})
	}

	user.ID = uuid.New().String()

	switch user.Tipo {
	case "Administrador":
		var admin entities.Administrador
		admin.Users = user
		admin.Permissao = true

	case "Professor":
		var professor entities.Professor
		professor.Users = user
		professor.Disciplina = "Disciplina"
		professor.Especialidade = "Especialidade"
		professor.Turno = "Turno"

	case "Aluno":
		var aluno entities.Aluno
		aluno.Users = user
		aluno.Serie = 1

	case "Responsavel":
		var responsavel entities.Responsavel
		responsavel.Users = user
		responsavel.Aluno = "ID Aluno"

	default:
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user type",
		})
	}

	return c.Status(http.StatusCreated).JSON(user)
}

//patch
func (uc *UsersController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var updatedUser entities.Users

	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse body",
		})
	}

	db := database.DB

	var existingUser entities.Users
	if err := db.First(&existingUser, "id = ?", id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	existingUser.Name = updatedUser.Name
	existingUser.Sobrenome = updatedUser.Sobrenome
	existingUser.Telefone = updatedUser.Telefone
	existingUser.Email = updatedUser.Email
	existingUser.Sexo = updatedUser.Sexo
	existingUser.CPF = updatedUser.CPF
	existingUser.Endereco = updatedUser.Endereco
	existingUser.DataNascimento = updatedUser.DataNascimento
	existingUser.Matricula = updatedUser.Matricula
	existingUser.Tipo = updatedUser.Tipo

	if err := db.Save(&existingUser).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.Status(http.StatusOK).JSON(existingUser)
}