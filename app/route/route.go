package route

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/ppmkh2sby/backend-user/app/controller"
	"github.com/ppmkh2sby/backend-user/app/repository"
	"github.com/ppmkh2sby/backend-user/app/usecase"
	"github.com/ppmkh2sby/backend-user/db"
)

// Init is function to initial route app
func Init(app *fiber.App, db *db.PostgresDB) {
	userRepository := repository.NewUserRepository(db.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: fmt.Sprintf("%s, %s, %s, %s", fiber.MethodGet, fiber.MethodPost, fiber.MethodPut, fiber.MethodDelete),
	}))

	api := app.Group("/api")
	api.Post("/signup", userController.SignUpUser)
	api.Post("/signin", userController.SignInUser)
	api.Get("/users", userController.GetAllUsers)
	api.Get("/users/:id", userController.GetUserByID)
	api.Put("/users/:id", userController.UpdateUser)
	api.Put("/users/:id", userController.ChangePasswordUser)
	api.Delete("/users/:id", userController.DeleteUser)
}
