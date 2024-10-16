package router

import (
	"fiber-mvc/app/common/constant"
	"fiber-mvc/app/controller"
	"fiber-mvc/app/service"
	"os"

	"github.com/damon35868/finalx-gofiber/common"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func ClientRegister(app fiber.Router, server *service.Service) {
	handler := controller.New(server)
	client := app.Group(constant.Client)

	clientUserApi := client.Group("/user")
	clientUserApi.Post("/login", handler.ClientController.Login)

	// JWT
	client.Use(jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(os.Getenv("JWT_CLIENT_SECRET"))},
		ErrorHandler: common.JWTErrorHandler,
	}))

	clientUserApi.Get("/", handler.ClientController.Info)
}