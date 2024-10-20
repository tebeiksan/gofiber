package exceptions

import "github.com/gofiber/fiber/v2"

func UserDeleteFailedException(c *fiber.Ctx, code int, messages ...string) error {

	message := "Delete user failed"

	if len(messages) == 0 {
		messages = append(messages, message)
	}

	return New(c, code, messages, GetFunctionName(UserDeleteFailedException))
}
