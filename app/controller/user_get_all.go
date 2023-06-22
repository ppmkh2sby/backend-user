package controller

import "github.com/gofiber/fiber/v2"

// GetAllUsers is function on controller to handle get all users
func (u *userController) GetAllUsers(c *fiber.Ctx) error {
	ctx := c.Context()

	users, err := u.userService.GetAllUsers(ctx)
	if err != nil {
		return err
	}

	return c.JSON(users)
}
