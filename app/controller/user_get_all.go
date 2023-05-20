package controller

import "github.com/gofiber/fiber/v2"

func (u *userController) GetAllUsers(c *fiber.Ctx) error {
	users, err := u.userService.GetAllUsers()
	if err != nil {
		return err
	}

	return c.JSON(users)
}
