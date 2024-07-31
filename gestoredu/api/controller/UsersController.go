package controller

import (
	"net/http"
	"regexp"
	"git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/api/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UsersController struct {}

func NewUsersController() *UsersController {
	return &UsersController{}
}


// As func de crate e de patch estao erradas
//-Melhorias a serem feitas, adicionar um falor especifico para cada tipo de user, ex ipotetico todo adm começa com 1
//com esse inicio de 1 as entities do amd passarim a ser o user tipo adm 
//patch seguiria essa mesma logica

//ou o user é criado, passamos o id desse user por um func que fara com que esse id receba as entities de adm por ex


func (u *UsersController) Create(c *fiber.Ctx) error {
	var baseUser entities.Users

	if err := c.BodyParser(&baseUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if !isValidBRDate(baseUser.DataNascimento) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid date format"})
	}

	baseUser.ID = uuid.New().String()

	var user entities.Users
	switch baseUser.Tipo {
	case "administrador":
		var administrador entities.Administrador
		if err := c.BodyParser(&administrador); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		administrador.Users = baseUser
		user = &administrador

	case "professor":
		var professor entities.Professor
		if err := c.BodyParser(&professor); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		professor.Users = baseUser
		user = &professor

	case "aluno":
		var aluno entities.Aluno
		if err := c.BodyParser(&aluno); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		aluno.Users = baseUser
		user = &aluno

	case "responsavel":
		var responsavel entities.Responsavel
		if err := c.BodyParser(&responsavel); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		responsavel.Users = baseUser
		user = &responsavel

	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user type"})
	}

	u.users = append(u.users, user)
	return c.Status(fiber.StatusOK).JSON(user)
}

func isValidBRDate(dateStr string) bool {
	re := regexp.MustCompile(`^\d{2}-\d{2}-\d{4}$`)
	return re.MatchString(dateStr)
}

func (u *UsersController) Patch(c *fiber.Ctx) error {
	id := c.Params("id")

	var foundUser interface{}
	for _, user := range u.users {
		switch v := user.(type) {
		case *entities.Administrador:
			if v.Users.ID == id {
				foundUser = v
			}
		case *entities.Professor:
			if v.Users.ID == id {
				foundUser = v
			}
		case *entities.Aluno:
			if v.Users.ID == id {
				foundUser = v
			}
		case *entities.Responsavel:
			if v.Users.ID == id {
				foundUser = v
			}
		}
	}

	if foundUser == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	switch user := foundUser.(type) {
	case *entities.Administrador:
		var updatedAdmin entities.Administrador
		if err := c.BodyParser(&updatedAdmin); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse JSON"})
		}
		updateBaseFields(&user.Users, updatedAdmin.Users)
		user.Permissao = updatedAdmin.Permissao
		foundUser = user

	case *entities.Professor:
		var updatedProfessor entities.Professor
		if err := c.BodyParser(&updatedProfessor); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse JSON"})
		}
		updateBaseFields(&user.Users, updatedProfessor.Users)
		user.Disciplina = updatedProfessor.Disciplina
		user.Especialidade = updatedProfessor.Especialidade
		user.Turno = updatedProfessor.Turno
		foundUser = user

	case *entities.Aluno:
		var updatedAluno entities.Aluno
		if err := c.BodyParser(&updatedAluno); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse JSON"})
		}
		updateBaseFields(&user.Users, updatedAluno.Users)
		user.Serie = updatedAluno.Serie
		foundUser = user

	case *entities.Responsavel:
		var updatedResponsavel entities.Responsavel
		if err := c.BodyParser(&updatedResponsavel); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse JSON"})
		}
		updateBaseFields(&user.Users, updatedResponsavel.Users)
		user.Aluno = updatedResponsavel.Aluno
		foundUser = user
	}

	return c.Status(fiber.StatusOK).JSON(foundUser)
}

func updateBaseFields(foundUser *entities.Users, updatedUser entities.Users) {
	if updatedUser.Name != "" {
		foundUser.Name = updatedUser.Name
	}
	if updatedUser.Sobrenome != "" {
		foundUser.Sobrenome = updatedUser.Sobrenome
	}
	if updatedUser.Telefone != 0 {
		foundUser.Telefone = updatedUser.Telefone
	}
	if updatedUser.Email != "" {
		foundUser.Email = updatedUser.Email
	}
	if updatedUser.Sexo != "" {
		foundUser.Sexo = updatedUser.Sexo
	}
	if updatedUser.Endereco != "" {
		foundUser.Endereco = updatedUser.Endereco
	}
	if updatedUser.Matricula != "" {
		foundUser.Matricula = updatedUser.Matricula
	}
	if updatedUser.Tipo != "" {
		foundUser.Tipo = updatedUser.Tipo
	}
	if updatedUser.CPF != 0 {
		foundUser.CPF = updatedUser.CPF
	}
	if isValidBRDate(updatedUser.DataNascimento) {
		foundUser.DataNascimento = updatedUser.DataNascimento
	}
}

//delete
func (t *UsersController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	for idx, user := range t.users {
		if user.ID == id {
			t.users = append(t.users[0:idx], t.users[idx+1:]...)
			return ctx.SendStatus(http.StatusOK)
			}
		}
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
	})
}