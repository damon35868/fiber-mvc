package router

import (
	"fiber-mvc/app/common/constant"
	"fiber-mvc/app/controller"
	"os"

	"github.com/damon35868/finalx-gofiber/common"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func ClientRegister(app fiber.Router, handler *controller.Controller) {
	client := app.Group(constant.Client)
	// 不走JWT鉴权
	client.Post("/user/login", handler.ClientController.Login)

	// JWT
	client.Use(jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(os.Getenv("JWT_CLIENT_SECRET"))},
		ErrorHandler: common.JWTErrorHandler,
	}))

	clientUserApi := client.Group("/user")
	{
		clientUserApi.Get("/info", handler.ClientController.Info)
	}
}
