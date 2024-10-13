package routes

import (
	"go_fiber_crud/app/controllers"
	"go_fiber_crud/configs"

	"github.com/gofiber/fiber/v2"
)

func Api(app *fiber.App, db *configs.Database) {
	controllers.UserController(app.Group("/user"), db)
}
