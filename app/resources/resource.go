package resources

import (
	"github.com/gofiber/fiber/v2"
)

func New(c *fiber.Ctx, data any, params ...int64) error {

	if len(params) > 0 {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "success",
			"data":    data,
			"total":   params[0],
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    data,
	})
}
