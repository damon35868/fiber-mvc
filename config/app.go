package config

import (
	"errors"

	"github.com/damon35868/finalx-gofiber/common"
	"github.com/gofiber/fiber/v2"
)

const (
	TIMER      = 1000
	INTER_TIME = 50
)

func NewApp() fiber.Config {
	return fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			return common.HttpException(ctx, code, err.Error())
		},
	}
}
