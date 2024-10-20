package exceptions

import "github.com/gofiber/fiber/v2"

func UserNotFoundException(c *fiber.Ctx, code int, messages ...string) error {

	message := "User not found"

	if len(messages) == 0 {
		messages = append(messages, message)
	}

	return New(c, code, messages)
}
