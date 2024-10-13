package main

import (
	"go_fiber_crud/configs"
	"go_fiber_crud/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	databaseConfig := configs.DatabaseConfig{
		Driver:   "mysql",
		Host:     "localhost",
		Username: "root",
		Password: "",
		Port:     3306,
		Database: "golang",
	}

	db, dbErr := configs.DatabaseNew(&databaseConfig)

	app := fiber.New()

	app.Use(recover.New())

	app.Use(func(c *fiber.Ctx) error {
		if dbErr != nil {
			return fiber.NewError(fiber.StatusServiceUnavailable, "Error DB connection")
		}

		return c.Next()
	})

	routes.Api(app, db)

	app.Listen(":3000")
}
