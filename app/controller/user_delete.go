package controller

import "github.com/gofiber/fiber/v2"

// DeleteUser is function on controller to handle delete user
func (u *userController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	err := u.userService.DeleteUser(id)
	if err != nil {
		return err
	}

	return c.JSON(err)
}
