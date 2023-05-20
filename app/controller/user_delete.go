package controller

import "github.com/gofiber/fiber/v2"

func (u *userController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	err := u.userService.DeleteUser(id)
	if err != nil {
		return err
	}

	return c.JSON(err)
}
