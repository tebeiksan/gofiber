package exceptions

import "github.com/gofiber/fiber/v2"

func UserUpdateFailedException(c *fiber.Ctx, code int, messages ...string) error {

	message := "Update user failed"

	if len(messages) == 0 {
		messages = append(messages, message)
	}

	return New(c, code, messages, GetFunctionName(UserUpdateFailedException))
}
