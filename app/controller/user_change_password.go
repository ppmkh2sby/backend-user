package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ppmkh2sby/backend-library/models"
)

// ChangePasswordUser is function on controller to handle change of user password
func (u *userController) ChangePasswordUser(c *fiber.Ctx) error {
	var user models.Users
	err := c.BodyParser(&user)
	if err != nil {
		return err
	}
	user.ID = c.Params("id")
	ctx := c.Context()

	err = u.userService.ChangePasswordUser(ctx, user.ID, user.Password)
	if err != nil {
		return err
	}

	return c.JSON(err)
}
