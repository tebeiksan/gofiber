package exceptions

import "github.com/gofiber/fiber/v2"

func BaseException(c *fiber.Ctx, code int, messages ...string) error {

	message := "Failed"

	if len(messages) == 0 {
		messages = append(messages, message)
	}

	return New(c, code, messages, GetFunctionName(BaseException))
}
