package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ppmkh2sby/backend-library/models"
)

// UpdateUser is controller to update data user on database
func (u *userController) UpdateUser(c *fiber.Ctx) error {
	var user models.Users
	err := c.BodyParser(&user)
	if err != nil {
		return err
	}
	user.ID = c.Params("id")

	response, err := u.userService.UpdateUser(&user)
	if err != nil {
		return err
	}

	return c.JSON(response)
}
