package middleware

import "github.com/gofiber/fiber/v2"

func Response(ctx *fiber.Ctx) error {
	return ctx.JSON(&fiber.Map{
		"code":    fiber.StatusOK,
		"message": "message",
	})
}
