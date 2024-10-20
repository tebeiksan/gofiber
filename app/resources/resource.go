package resources

import "github.com/gofiber/fiber/v2"

func New(c *fiber.Ctx, data any) error {
	return c.JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    data,
	})
}
