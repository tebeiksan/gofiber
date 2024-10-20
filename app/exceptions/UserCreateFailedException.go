package exceptions

import "github.com/gofiber/fiber/v2"

func UserCreateFailedException(c *fiber.Ctx, code int, messages ...string) error {

	message := "Create user failed"

	if len(messages) == 0 {
		messages = append(messages, message)
	}

	return New(c, code, messages, GetFunctionName(UserCreateFailedException))
}
