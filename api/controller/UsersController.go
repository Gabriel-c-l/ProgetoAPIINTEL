package controller

import (
	"net/http"
	"github.com/gofiber/fiber/v2"
	"git.intelbras.com.br/isec/linha-future/jovens/gestoredu.git/models"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type UsersController struct {
	DB *gorm.DB
}

func NewUsersController(db *gorm.DB) *UsersController {
	return &UsersController{DB: db}
}

// CreateUser cria um novo usuário
func (uc *UsersController) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse body",
		})
	}

	user.ID = uuid.New().String()
	if err := uc.DB.Create(&user).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// GetAllUsers retorna todos os usuários
func (uc *UsersController) GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	if err := uc.DB.Find(&users).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}

	return c.Status(http.StatusOK).JSON(users)
}

// GetUserByID retorna um usuário pelo ID
func (uc *UsersController) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var user models.User
	if err := uc.DB.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch user",
		})
	}

	return c.Status(http.StatusOK).JSON(user)
}

// DeleteUserByID exclui um usuário pelo ID
func (uc *UsersController) DeleteUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := uc.DB.Delete(&models.User{}, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	return c.SendStatus(http.StatusNoContent)
}

// UpdateUserByID atualiza um usuário pelo ID
func (uc *UsersController) UpdateUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var updates models.User
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	var user models.User
	if err := uc.DB.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch user",
		})
	}

	if err := uc.DB.Model(&user).Updates(updates).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}

	return c.Status(http.StatusOK).JSON(user)
}
