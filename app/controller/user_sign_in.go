package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ppmkh2sby/backend-library/models"
)

// SignInUser is controller to sign user
func (u *userController) SignInUser(c *fiber.Ctx) error {
	var user models.Users
	err := c.BodyParser(&user)
	if err != nil {
		return err
	}

	response, err := u.userService.SignInUser(&user)
	if err != nil {
		return err
	}

	return c.JSON(response)
}
