package controllers

import (
	"go_fiber_crud/app/helpers"
	"go_fiber_crud/app/models"
	"go_fiber_crud/configs"

	"github.com/gofiber/fiber/v2"
)

func UserController(route fiber.Router, db *configs.Database) {
	route.Get("/", UserIndex(db))
	route.Post("/", UserCreate)
}

func UserIndex(db *configs.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var User []models.User
		var UserCount int64

		db.Model(&User).Count(&UserCount)

		if response := db.Scopes(helpers.Paginate(c)).Find(&User); response.Error != nil {
			panic("Error occurred while retrieving roles from the database: " + response.Error.Error())
		}
		err := c.JSON(User)

		if err != nil {
			panic("Error occurred when returning JSON of user: " + err.Error())
		}

		return c.JSON(fiber.Map{
			"success":    true,
			"usersTotal": UserCount,
			"users":      User,
		})
	}
}

func UserCreate(c *fiber.Ctx) error {
	return c.SendString("User Create")
}
