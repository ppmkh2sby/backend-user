package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ppmkh2sby/backend-library/models"
)

// SignUpUser is function on controller to handle sign up user
func (u *userController) SignUpUser(c *fiber.Ctx) error {
	var user models.Users
	err := c.BodyParser(&user)
	if err != nil {
		return err
	}
	ctx := c.Context()

	response, err := u.userService.SignUpUser(ctx, &user)
	if err != nil {
		return err
	}

	return c.JSON(response)
}
