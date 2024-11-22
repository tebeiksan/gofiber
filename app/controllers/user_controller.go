package controllers

import (
	"errors"
	"fmt"
	"go_fiber_crud/app/exceptions"
	"go_fiber_crud/app/helpers"
	"go_fiber_crud/app/models"
	"go_fiber_crud/app/resources"
	"go_fiber_crud/configs"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UserController(route fiber.Router, db *configs.Database) {
	route.Get("/", UserIndex(db))
	route.Post("/", UserCreate(db))
	route.Get("/:id", UserShow(db))
	route.Put("/:id", UserUpdate(db))
	route.Delete("/:id", UserDelete(db))
}

func UserIndex(db *configs.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var User []models.User
		var UserCount int64

		// getUsers := db.Scopes(helpers.Paginate(c)).Where("email like ?", search).Find(&User)
		// getUsers := db.Where("email like ?", search).Session(&gorm.Session{})

		getUsers := db.Model(&User)

		var search = c.Query("search")

		if strings.TrimSpace(search) != "" {
			getUsers.Where("email like ?", "%"+search+"%").Session(&gorm.Session{})
		}

		getUsers.Count(&UserCount)
		getUsers.Scopes(helpers.Paginate(c)).Find(&User)

		if getUsers.Error != nil {
			return exceptions.DatabaseException(c, fiber.ErrInternalServerError.Code, fmt.Sprint(getUsers.Error))
		}

		return resources.New(c, User, UserCount)

	}
}

func UserCreate(db *configs.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := new(models.User)
		return exceptions.UserCreateFailedException(c, fiber.ErrInternalServerError.Code, fmt.Sprint("error coba cok"))

		if err := c.BodyParser(request); err != nil {
			return err
		}

		uuid := uuid.New().String()

		user := models.User{Id: uuid, Name: request.Name, Email: request.Email}

		createUser := db.Create(&user)

		if createUser.Error != nil {
			return exceptions.UserCreateFailedException(c, fiber.ErrInternalServerError.Code, fmt.Sprint(createUser.Error))
		}

		return resources.New(c, user)

	}
}

func UserShow(db *configs.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user = models.User{Id: c.Params("id")}

		getUser := db.First(&user)

		if errors.Is(getUser.Error, gorm.ErrRecordNotFound) {
			return exceptions.UserNotFoundException(c, fiber.ErrNotFound.Code)
		}

		return resources.New(c, user)
	}
}

func UserUpdate(db *configs.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user = models.User{Id: c.Params("id")}

		getUser := db.First(&user)

		if errors.Is(getUser.Error, gorm.ErrRecordNotFound) {
			return exceptions.UserNotFoundException(c, fiber.ErrNotFound.Code)
		}

		request := new(models.User)

		if err := c.BodyParser(request); err != nil {
			return exceptions.BaseException(c, fiber.ErrUnprocessableEntity.Code, fmt.Sprint(err))
		}

		updateUser := db.Model(&user).Updates(models.User{Name: request.Name})

		if updateUser.Error != nil {
			return exceptions.UserUpdateFailedException(c, fiber.ErrUnprocessableEntity.Code, fmt.Sprint(updateUser.Error))
		}

		return resources.New(c, user)

	}
}

func UserDelete(db *configs.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user = models.User{Id: c.Params("id")}

		getUser := db.First(&user)

		if errors.Is(getUser.Error, gorm.ErrRecordNotFound) {
			return exceptions.UserNotFoundException(c, fiber.ErrNotFound.Code)
		}

		deleteUser := db.Delete(&user)

		if deleteUser.Error != nil {
			return exceptions.UserDeleteFailedException(c, fiber.ErrNotFound.Code, fmt.Sprint(deleteUser.Error))
		}

		return resources.New(c, user)
	}
}
