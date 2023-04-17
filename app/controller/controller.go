package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ppmkh2sby/backend-user/app/usecase"
)

type userController struct {
	userService usecase.UserUsecase
}

type UserController interface {
	SignUpUser(c fiber.Ctx) error
	SingInUser(c fiber.Ctx) error
	GetAllUsers(c fiber.Ctx) error
	GetUserByID(c fiber.Ctx) error
	UpdateUser(c fiber.Ctx) error
	DeleteUser(c fiber.Ctx) error
	ChangePasswordUser(c fiber.Ctx) error
}

// NewUserController is function to intial User Controller
func NewUserController(u usecase.UserUsecase) *userController {
	return &userController{
		userService: u,
	}
}
