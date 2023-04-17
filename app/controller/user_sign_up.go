package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ppmkh2sby/backend-library/models"
)

func (u *userController) SignUpUser(c fiber.Ctx) error {
	var user models.Users
	err := c.BodyParser(&user)
	if err != nil {
		return err
	}

	response, err := u.userService.SignUpUser(&user)
	if err != nil {
		return err
	}

	return c.JSON(response)
}
