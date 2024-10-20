package exceptions

import "github.com/gofiber/fiber/v2"

func New(c *fiber.Ctx, code int, messages []string) error {

	var message string

	for _, v := range messages {
		message = v
		break
	}

	if code == 0 {
		code = fiber.ErrServiceUnavailable.Code
	}

	c.SendStatus(code)

	return c.JSON(fiber.Map{
		"success": false,
		"message": message,
	})
}
