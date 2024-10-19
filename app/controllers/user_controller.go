package controllers

import (
	"errors"
	"fmt"
	"go_fiber_crud/app/helpers"
	"go_fiber_crud/app/models"
	"go_fiber_crud/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UserController(route fiber.Router, db *configs.Database) {
	route.Get("/", UserIndex(db))
	route.Post("/", UserCreate(db))
	route.Get("/:Id", UserShow(db))
	route.Put("/:Id", UserUpdate(db))
	route.Delete("/:Id", UserDelete(db))
}

func UserIndex(db *configs.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var User []models.User
		var UserCount int64

		db.Model(&User).Count(&UserCount)

		var search = "%" + c.Query("search") + "%"

		if response := db.Scopes(helpers.Paginate(c)).Where("email like ?", search).Find(&User); response.Error != nil {
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

func UserCreate(db *configs.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(models.User)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		uuid := uuid.New().String()

		user := models.User{Id: uuid, Name: request.Name, Email: request.Email}

		result := db.Create(&user)

		if result.Error != nil {
			return c.JSON(fiber.Map{
				"success": false,
				"message": fmt.Sprint(result.Error),
			})
		}

		return c.JSON(fiber.Map{
			"success": true,
			"users":   user,
		})

	}
}

func UserShow(db *configs.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user = models.User{Id: c.Params("Id")}

		result := db.First(&user)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(fiber.Map{
				"success": false,
				"message": "User not found",
			})
		}

		return c.JSON(fiber.Map{
			"success": true,
			"users":   user,
		})
	}
}

func UserUpdate(db *configs.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user = models.User{Id: c.Params("Id")}

		result := db.First(&user)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(fiber.Map{
				"success": false,
				"message": "User not found",
			})
		}

		request := new(models.User)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		db.Model(&user).Updates(models.User{Name: request.Name})

		return c.JSON(fiber.Map{
			"success": true,
			"users":   user,
		})
	}
}

func UserDelete(db *configs.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user = models.User{Id: c.Params("Id")}

		result := db.First(&user)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.JSON(fiber.Map{
				"success": false,
				"message": "User not found",
			})
		}

		db.Delete(&user)

		return c.JSON(fiber.Map{
			"success": true,
			"users":   user,
		})
	}
}
