package main

import (
	"go_fiber_crud/app/exceptions"
	"go_fiber_crud/configs"
	"go_fiber_crud/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	databaseConfig := configs.DatabaseConfig{
		Driver:   "mysql",
		Host:     "localhost",
		Username: "root",
		Password: "",
		Port:     3306,
		Database: "services",
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Use(recover.New())

	db, dbErr := configs.DatabaseNew(&databaseConfig)

	app.Use(func(c *fiber.Ctx) error {
		if dbErr != nil {
			return exceptions.DatabaseException(c, fiber.ErrServiceUnavailable.Code)
		}

		return c.Next()
	})

	routes.Api(app, db)

	app.Listen(":3333")
}
