package controller

import "github.com/gofiber/fiber/v2"

// GetUserByID is function on controller to handle get specific user by id
func (u *userController) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := u.userService.GetUserByID(id)
	if err != nil {
		return err
	}

	return c.JSON(user)
}
