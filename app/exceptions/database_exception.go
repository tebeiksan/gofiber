package exceptions

import "github.com/gofiber/fiber/v2"

func DatabaseException(c *fiber.Ctx, code int, messages ...string) error {

	message := "Error database connection"

	if len(messages) == 0 {
		messages = append(messages, message)
	}

	return New(c, code, messages)
}
