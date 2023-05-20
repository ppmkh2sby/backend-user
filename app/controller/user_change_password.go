package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ppmkh2sby/backend-library/models"
)

func (u *userController) ChangePasswordUser(c *fiber.Ctx) error {
	var user models.Users
	err := c.BodyParser(&user)
	if err != nil {
		return err
	}
	user.ID = c.Params("id")

	err = u.userService.ChangePasswordUser(user.ID, user.Password)
	if err != nil {
		return err
	}

	return c.JSON(err)
}
