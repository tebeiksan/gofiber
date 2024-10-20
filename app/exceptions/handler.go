package exceptions

import (
	"reflect"
	"runtime"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetFunctionName(i interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func New(c *fiber.Ctx, code int, messages []string, exceptionName string) error {

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
		"data":    map[string]string{"exception": exceptionName},
	})
}
